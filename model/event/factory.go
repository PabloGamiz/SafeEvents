package event

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
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

// OpenEventStream opens an stream ensuring the client's table does exists
func OpenEventStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got an error while opening stream: %v", err.Error())
		return
	}

	once.Do(func() {
		db.AutoMigrate(&service.Service{}, &Event{})
	})
	return
}

//
//func LoadOrStoreNewEvent(event Controller) error {
//	sid := event.GetID()
//	if
//}

// FindAll returns the controllers of all the events loaded on the BBDD
func FindAll(ctx context.Context) (ctrl []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}

	var eventsMOD []*Event
	db.Preload("Services.Location").Preload("Services.Products").Preload(clause.Associations).Find(&eventsMOD)

	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}
	var events []*Event
	db.Preload("Services.Location").Preload("Services.Products").Preload(clause.Associations).Where("id = ?", ID).Find(&events)
	if len(events) == 0 {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}

	ctrl = events[0]

	return
}
