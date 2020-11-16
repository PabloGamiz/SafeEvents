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
	mu    sync.Mutex
	touch bool = false
)

// OpenClientStream opens an stream ensuring the client's table does exists
func OpenClientStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while opening stream", err.Error())
		return
	}

	if !touch {
		mu.Lock()
		defer mu.Unlock()
		if !touch {
			db.AutoMigrate(&client.Client{})
			touch = true
		}
	}

	return
}

// NewClientGateway builds a gateway for the provided client
func NewClientGateway(ctx context.Context, client *client.Client) Gateway {
	return &clientGateway{Client: client, ctx: ctx}
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
