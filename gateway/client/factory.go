package client

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
)

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
		db.AutoMigrate(&client.Client{})
	})

	return
}

// NewClientGateway builds a gateway for the provided client
func NewClientGateway(ctx context.Context, client client.Controller) Gateway {
	return &clientGateway{Controller: client, ctx: ctx}
}

// FindClientByEmail returns the gateway for the client that match the provided mail
func FindClientByEmail(ctx context.Context, email string) (gw Gateway, err error) {
	var db *gorm.DB
	if db, err = OpenClientStream(); err != nil {
		return
	}

	var clients []client.Client
	db.Where(queryFindByEmail, email).Find(&clients)
	if len(clients) == 0 {
		err = fmt.Errorf(errNotFoundByEmail, email)
		return
	}

	gw = NewClientGateway(ctx, &clients[0])
	return
}

// FindClientByID returns the gateway for the client that match the provided mail
func FindClientByID(ctx context.Context, ID uint) (gw Gateway, err error) {
	var db *gorm.DB
	if db, err = OpenClientStream(); err != nil {
		return
	}

	var client client.Client
	notR := db.Where(queryFindByID, ID).Find(&client)
	if notR != nil {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}

	gw = NewClientGateway(ctx, &client)
	return
}
