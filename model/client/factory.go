package client

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
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

	var clients []*Client
	if result := db.Where(queryFindByEmail, email).Find(&clients); result.Error != nil {
		err = fmt.Errorf(errNotFoundByEmail, result.Error.Error(), email)
		return
	}

	client := clients[0]
	if assistant := client.Assists; assistant != nil {
		assistant.SetParent(client)
	}

	if organizer := client.Organize; organizer != nil {
		organizer.SetParent(client)
	}

	return client, nil
}

// FindClientByID returns the gateway for the client that match the provided mail
func FindClientByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenClientStream(); err != nil {
		return
	}

	var clients []*Client
	if result := db.Where(queryFindByID, ID).Find(&clients); result.Error != nil {
		err = fmt.Errorf(errNotFoundByID, result.Error.Error(), ID)
		return
	}

	client := clients[0]
	if assistant := client.Assists; assistant != nil {
		assistant.SetParent(client)
	}

	if organizer := client.Organize; organizer != nil {
		organizer.SetParent(client)
	}

	return client, nil
}
