package event

import (
	"context"
	"fmt"
	"sync"

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

//
//func LoadOrStoreNewEvent(event Controller) error {
//	sid := event.GetID()
//	if
//}

// FindAll returns the controllers of all the events loaded on the BBDD
func FindAll(ctx context.Context) (ctrl []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var eventsMOD []*Event
	db.Preload(clause.Associations).Preload("Services.Location").Preload("Services.Products").Find(&eventsMOD)

	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
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
