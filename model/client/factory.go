package client

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"

	"gorm.io/gorm"
)

// FindClientByEmail returns the gateway for the client that match the provided mail
func FindClientByEmail(ctx context.Context, email string) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	var client Client
	//.Preload("Assists.Purchased").Preload("Organize.Organize")
	db = db.Where(queryFindByEmail, email).Find(&client)

	if db.Error != nil {
		err = fmt.Errorf(errNotFoundByEmail, db.Error.Error(), email)
		return
	}

	if client.GetID() == 0 {
		return
	}

	return &client, nil
}

// FindClientByID returns the gateway for the client that match the provided mail
func FindClientByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	var client Client
	if db = db.Preload("Assists.Purchased").Preload("Organize.Organize").Where(queryFindByID, ID).Find(&client); db.Error != nil {
		err = fmt.Errorf(errNotFoundByID, db.Error.Error(), ID)
		return
	}

	if client.GetID() == 0 {
		err = fmt.Errorf(errNotFoundByID, "no value", ID)
		return
	}

	return &client, nil
}

// // AddOrganizer ...
// func AddOrganizer(ctx context.Context, event event.Controller, clt Controller) (ctrl Controller, err error) {
// 	return
// }

// FindOrganitzersEvent returns the gateways for the clients that organize the provided email
func FindOrganitzersEvent(ctx context.Context, EventID uint) (NameOrg string, err error) {
	/*var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}*/
	//NO FA RES
	return
}

// FindAllFavs returns the gateway for finding all the events loaded on the BBDD
func FindAllFavs(ctx context.Context, clientCtrl Controller) (events []event.Controller, err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	var eventsMOD []*event.Event
	db.Model(clientCtrl).Association("Favs").Find(&eventsMOD)
	if len(eventsMOD) == 0 {
		err = fmt.Errorf("errNoEventsOnDatabase")
		return
	}

	events = make([]event.Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		events[index] = event
	}

	return
}
