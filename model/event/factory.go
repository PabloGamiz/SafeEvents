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
	if db, err = mysql.OpenStream(); err != nil {
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
func FindEventByID(ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
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
