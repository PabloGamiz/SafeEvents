package event

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

var (
	mu    sync.Mutex
	touch bool = false
)

// OpenEventStream opens an stream ensuring the client's table does exists
func OpenEventStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got an error while opening stream: %v", err.Error())
		return
	}

	if !touch {
		mu.Lock()
		defer mu.Unlock()
		if !touch {
			db.AutoMigrate(&event.Event{})
			touch = true
		}
	}

	return
}

// NewEventGateway builds a gateway for the provided event
func NewEventGateway(ctx context.Context, event event.Controller) Gateway {
	return &eventGateway{Controller: event, ctx: ctx}
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ctx context.Context, ID int) (gw Gateway, err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}
	var events []event.Event
	db.Where("id = ?", ID).Find(&events)
	if len(events) == 0 {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}

	gw = NewEventGateway(ctx, &events[0])

	return
}
