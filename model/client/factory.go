package client

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"

	"gorm.io/gorm"
)

var once sync.Once

// OpenClientStream opens an stream ensuring the client's table does exists
func OpenClientStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while opening stream", err.Error())
		return
	}

	once.Do(func() {
		// Automigrate must be called only once for each gateway, and allways on the stream's opening call.
		// This makes sure the client struct has its own table on the database. So model updates are only
		// migrable to the database rebooting the server (not on-the-run).
		db.AutoMigrate(&Client{}, &organizer.Organizer{}, &assistant.Assistant{})
	})

	return
}

// FindClientByEmail returns the gateway for the client that match the provided mail
func FindClientByEmail(ctx context.Context, email string) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenClientStream(); err != nil {
		return
	}

	var client Client
	if result := db.Where(queryFindByEmail, email).Find(&client); result.Error != nil {
		err = fmt.Errorf(errNotFoundByEmail, result.Error.Error(), email)
		return
	}

	if client.GetID() == 0 {
		err = fmt.Errorf(errNotFoundByEmail, "no value", email)
		return
	}

	client.GetAssistant().SetParent(client)
	client.GetOrganizer().SetParent(client)

	return &client, nil
}

// FindClientByID returns the gateway for the client that match the provided mail
func FindClientByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenClientStream(); err != nil {
		return
	}

	var client Client
	if db = db.Where(queryFindByID, ID).Find(&client); db.Error != nil {
		err = fmt.Errorf(errNotFoundByID, db.Error.Error(), ID)
		return
	}

	if client.GetID() == 0 {
		err = fmt.Errorf(errNotFoundByID, "no value", ID)
		return
	}

	client.GetAssistant().SetParent(client)
	client.GetOrganizer().SetParent(client)

	return &client, nil
}

// AddOrganizer ...
func AddOrganizer(ctx context.Context, evnt event.Controller, clt Controller) (ctrl Controller, err error) {
	return
}
