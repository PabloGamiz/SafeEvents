package event

import (
	"context"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
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

// FindAll returns the controllers of all the events loaded on the BBDD
func FindAll(ctx context.Context) (ctrl []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var eventsMOD []*Event
	timeNow := time.Now()
	db.Preload(clause.Associations).Where(queryFindAll, timeNow).Find(&eventsMOD)
	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindAllByType returns the controllers of all the events loaded on the BBDD
func FindAllByType(ctx context.Context, eventType string) (ctrl []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()

	var eventsMOD []*Event
	timeNow := time.Now()
	db.Preload(clause.Associations).Where(queryFilterByType, eventType, timeNow).Find(&eventsMOD)
	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}
	var events []*Event
	defer cancel()
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
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}
	m := map[string]int{}
	m["Art"] += 5
	var PopularEvts []*Event
	defer cancel()
	var idAssis uint
	db.Select("id").Where("client_id = ?", clientID).Table("assistants").Find(&idAssis)
	//db.Table("events").Where("capacity <> ?", "taken").Where("closure_date > ?", time.Now()).Find(&PopularEvts).Order("taken desc")
	db.Raw("SELECT * FROM `events` WHERE capacity <> taken AND closure_date > CURRENT_DATE ORDER BY taken desc").Scan(&PopularEvts) //OBTE TOTS ELS ESDEVENIMENTS DISPONIBLES ORDENATS PER PLACES LLIURES
	type Result struct {
		EventID uint
	}
	var result []Result
	//GUARDA LES ASOCIACIONS TICKET_ID I ASSISTANT_ID DELS ESDEVENIMENTS PASATS
	db.Raw("SELECT DISTINCT tickets.event_id FROM tickets INNER JOIN events on tickets.event_id = events.id WHERE events.closure_date < CURRENT_DATE AND tickets.assistant_id = ?", idAssis).Scan(&result)
	if len(PopularEvts) == 0 {
		err = fmt.Errorf("Not Available Events for client ID %d", clientID)
		return
	}
	//CONSULTA ELS RATINGS DELS FEEDBACKS
	for _, valor := range result {
		var fb feedback.Feedback
		db.Preload(clause.Associations).Table("feedbacks").Where("event_id = ?", valor.EventID).Where("assistant_id = ?", idAssis).Find(&fb)
		var evnt []*Event
		db.Table("events").Where("id = ?", valor.EventID).Find(&evnt)
		if len(evnt) != 0 {
			if fb.ID != 0 {
				if fb.Rating > 3 {
					m[evnt[0].Tipus] += 10
				} else {
					m[evnt[0].Tipus] -= 10
				}
			} else {
				m[evnt[0].Tipus] += 5
			}
		}
	}
	m["Esports"] += 11
	//CONSULTA ELS FAVS DEL CLIENT
	var event_id []uint
	db.Raw("SELECT `event_id` FROM `clients_favs` WHERE client_id = ?", clientID).Scan(&event_id)
	if len(event_id) != 0 {
		for _, valor := range event_id {
			var evnt []*Event
			db.Table("events").Where("id = ?", valor).Find(&evnt)
			if len(evnt) != 0 {
				m[evnt[0].Tipus] += 10
			}
		}
	}
	for i, valor := range m {
		if valor < 0 {
			delete(m, i)
		}
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, m[k])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	fmt.Println(m, keys)
	resultsize := int(math.Min(10, float64(len(PopularEvts))))
	ctrl = make([]Controller, resultsize)
	//mapiter := 0
	for i := 0; i < resultsize; i++ {
		ctrl[i] = PopularEvts[i]
	}
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
