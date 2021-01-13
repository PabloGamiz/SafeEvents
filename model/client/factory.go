package client

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// FindClientByEmail returns the gateway for the client that match the provided mail
func FindClientByEmail(ctx context.Context, email string) (ctrl Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var client Client
	if result := db.Preload(clause.Associations).Preload("Assists.Purchased").Preload("Organize.Events").Where(queryFindByEmail, email).Find(&client); result.Error != nil {
		err = fmt.Errorf(errNotFoundByEmail, result.Error.Error(), email)
		return
	}

	if client.GetID() == 0 {
		err = fmt.Errorf(errNotFoundByEmail, "no value", email)
		return
	}

	client.GetAssistant().SetParent(&client)
	client.GetOrganizer().SetParent(&client)

	return &client, nil
}

// FindClientByID returns the gateway for the client that match the provided mail
func FindClientByID(ctx context.Context, ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var client Client
	if db = db.Preload(clause.Associations).Preload("Assists.Purchased").Preload("Organize.Events").Where(queryFindByID, ID).Find(&client); db.Error != nil {
		err = fmt.Errorf(errNotFoundByID, db.Error.Error(), ID)
		return
	}

	if got := client.GetID(); got != ID {
		err = fmt.Errorf(errNotFoundByID, got, ID)
		return
	}

	client.GetAssistant().SetParent(&client)
	client.GetOrganizer().SetParent(&client)

	return &client, nil
}

// // AddOrganizer ...
// func AddOrganizer(ctx context.Context, event event.Controller, clt Controller) (ctrl Controller, err error) {
// 	return
// }

// FindOrganitzersEvent returns the gateways for the clients that organize the provided email
func FindOrganitzersEvent(ctx context.Context, EventID uint) (NameOrg string, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	//var clientsMOD []*organizer.Organizer
	var orgIDs []uint
	var cltIDs []uint
	//var eventsMOD []*event.Event
	//var eventMOD *event.Event
	//subQuery := db.Table("events").Where("id = ?", EventID).Find(&eventMOD)
	//db.Table("organizers_events").Preload(clause.Associations).Where("event_id = ?", EventID).Find(&clientsMOD)
	db.Raw("SELECT organizer_id FROM organizers_events WHERE event_id = ?", EventID).Scan(&orgIDs)
	if len(orgIDs) == 0 {
		err = fmt.Errorf("errNoEventsOnDatabase")
		return
	}
	db.Raw("SELECT client_id FROM organizers WHERE id = ?", orgIDs[0]).Scan(&cltIDs)
	if len(cltIDs) == 0 {
		err = fmt.Errorf("errNoEventsOnDatabase")
		return
	}
	cltController, err := FindClientByID(ctx, cltIDs[0])
	NameOrg = cltController.GetEmail()
	return
}

// FindAllFavs returns the gateway for finding all the events loaded on the BBDD
func FindAllFavs(ctx context.Context, clientCtrl Controller) (events []event.Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var eventsMOD []*event.Event
	db.Model(clientCtrl).Association("Favs").Find(&eventsMOD)
	if len(eventsMOD) == 0 {
		//err = fmt.Errorf("errNoEventsOnDatabase")
		return make([]event.Controller, 0), err
	}

	events = make([]event.Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		events[index] = event
	}

	return
}

// FindClientEmailByClientID returns the email of the client matching the given ID.
func FindClientEmailByClientID(ctx context.Context, ID uint) (email string, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	if db = db.Select("email").Where(queryFindByID, ID).Table("clients").Find(&email); db.Error != nil {
		err = fmt.Errorf(errNotFoundByID, db.Error.Error(), ID)
		return
	}

	return
}
