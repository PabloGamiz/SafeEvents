package event

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	once sync.Once
)

// OpenEventStream opens an stream ensuring the client's table does exists
func OpenEventStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got an error while opening stream: %v", err.Error())
		return
	}

	once.Do(func() {
		db.AutoMigrate(&event.Event{})
	})
	return
}

// NewEventGateway builds a gateway for the provided client
func NewEventGateway(ctx context.Context, event *event.Event) Gateway {
	return &eventGateway{Event: event, ctx: ctx}
}

// FindAll returns the gateway for finding all the events loaded on the BBDD
func FindAll() (events []event.Event, err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}

	db.Preload("Services.Location").Preload("Services.Products").Preload(clause.Associations).Find(&events)
	if len(events) == 0 {
		err = fmt.Errorf(errNoEventsOnDatabase)
		return
	}

	return
}
