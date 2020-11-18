package event

import (
	"context"
	"fmt"
	"log"
	"os/user"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var once sync.Once

// OpenEventStream opens an stream ensuring the client's table does exists
func OpenEventStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got an error while opening stream: %v", err.Error())
		return
	}

	once.Do(func() {
		db.AutoMigrate(&event.Event{}, &user.User{})
	})
	return
}

// NewEventGateway builds a gateway for the provided event
func NewEventGateway(ctx context.Context, event event.Controller) Gateway {
	return &eventGateway{Controller: event, ctx: ctx}
}

// FindAll returns the gateway for finding all the events loaded on the BBDD
func FindAll(ctx context.Context) (events []event.Controller, err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}

	var eventsMOD []*event.Event
	db.Preload("Services.Location").Preload("Services.Products").Preload(clause.Associations).Find(&eventsMOD)
	if len(eventsMOD) == 0 {
		err = fmt.Errorf(errNoEventsOnDatabase)
		return
	}

	events = make([]event.Controller, len(eventsMOD))
	fmt.Println(events)
	for index, event := range eventsMOD {
		events[index] = NewEventGateway(ctx, event)
	}

	return
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ctx context.Context, ID int) (gw Gateway, err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}
	var events []event.Event
	db.Preload("Services.Location").Preload("Services.Products").Preload(clause.Associations).Where("id = ?", ID).Find(&events)
	if len(events) == 0 {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}

	gw = NewEventGateway(ctx, &events[0])

	return
}
