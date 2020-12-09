package event

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	// AllInstancesByID stores all events indexed by its ID
	AllInstancesByID = &sync.Map{}
	once             sync.Once
)

type sID uint

// OpenStream ensuring the client's table does exists
func OpenStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got an error while OpenStream: %v", err.Error())
		return
	}

	once.Do(func() {
		db.AutoMigrate(&service.Service{}, &Event{})
	})
	return
}

// FindAll returns the controllers of all the events loaded on the BBDD
func FindAll(ctx context.Context) (ctrl []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenStream(); err != nil {
		return
	}

	var eventsMOD []*Event
	db.Preload(clause.Associations).Find(&eventsMOD)
	fmt.Println(eventsMOD)
	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindAllByType returns the controllers of all the events loaded on the BBDD
func FindAllByType(ctx context.Context, eventType string) (ctrl []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenStream(); err != nil {
		return
	}

	var eventsMOD []*Event
	db.Preload(clause.Associations).Where(queryFilterByType, eventType).Find(&eventsMOD)
	fmt.Println(eventsMOD)
	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenStream(); err != nil {
		return
	}
	var events []*Event
	db.Preload(clause.Associations).Preload("Services.Products").Where("id = ?", ID).Find(&events)
	if len(events) == 0 {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}

	ctrl = events[0]

	return
}

// FindRecomended returns the controllers of all the events recomended for the user
func FindRecomended(ctx context.Context, clientID uint) (ctrl []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenStream(); err != nil {
		return
	}
	m := map[string]int{
		"Musica":  0,
		"Teatre":  0,
		"Esports": 0,
		"Art":     0,
	}
	var t []*ticket.Ticket
	var PopularEvts []*Event
	idAssis := db.Select("id").Where("client_id = ?", clientID).Table("assistants")
	db.Table("events").Where("capacity <> ?", "taken").Where("closure_date > ?", time.Now()).Find(&PopularEvts).Order("taken desc").Limit(10)
	subquery := db.Table("events").Where("closure_date < ?", time.Now())
	db.Table("tickets").Where("assistant_id = ?", idAssis).Where("event_id = (?)", subquery).Group("event_id").Group("assistant_id").Find(&t)

	if len(PopularEvts) == 0 {
		err = fmt.Errorf("Not Available Events %d", clientID)
		return
	}
	for _, valor := range t {
		var fb feedback.Feedback
		db.Table("feedbacks").Where("event_id = ?", valor.EventID).Where("assistant_id = ?", idAssis).Find(&fb)
		if fb.ID != 0 {
			var evnt *Event
			db.Table("events").Where("event_id = ?", valor.EventID).Find(&evnt)
			if fb.Rating > 3 {
				m[evnt.Tipus] += 10
			} else {
				m[evnt.Tipus] -= 10
			}
		}
	}
	resultsize := int(math.Min(10, float64(len(PopularEvts))))
	ctrl = make([]Controller, resultsize)
	/*	if len(t) == 0 {
			err = fmt.Errorf("Not enough data from the user", clientID)
			//devolver eventos con mayor popularidad (menos tickets disponibles)
			//return
		}

		for index, valor := range m {
			rows, _ := db.Table("events").Select("id").Where("tipus", index).Rows()
			defer rows.Close()

			for rows.Next() {
				var evnt Event
				var count int64
				db.ScanRows(rows, &evnt)
				db.Table("tickets").Group("event_id").Where("assistant_id = ?", idAssis).Where("event_id = ?", evnt.ID).Count(&count)

				//find ratings at feedbacks
				var fb feedback.Feedback
				db.Table("feedbacks").Where("event_id = ?", evnt.ID).Where("assistant_id = ?", idAssis).Find(&fb)
				m[index] = int(count) + valor
			}
		}
	*/
	return
}
