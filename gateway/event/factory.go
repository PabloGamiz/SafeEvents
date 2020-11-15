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

<<<<<<< HEAD
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
=======
var (
	once sync.Once
>>>>>>> feature/event-management/publica_i_consulta_especifica
)

// OpenEventStream opens an stream ensuring the Event's table does exists
func OpenEventStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while opening stream", err.Error())
		return
	}

	once.Do(func() {
		// Automigrate must be called only once for each gateway, and allways on the stream's opening call.
		// This makes sure the Event struct has its own table on the database. So model updates are only
		// migrable to the database rebooting the server (not on-the-run).
		db.AutoMigrate(&event.Event{})
	})

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
	var event event.Event
	notR := db.Where("ID = ?", ID).Find(&event)
	if notR != nil {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}
	return
}
