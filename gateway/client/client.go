package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"gorm.io/gorm"
)

type clientGateway struct {
	client.Controller
	ctx context.Context
}

func (gw *clientGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	return db.Create(gw.Controller).Error
}

func (gw *clientGateway) Update() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	if db = db.Save(gw.Controller); db.Error != nil {
		return db.Error
	}
	/*
			for _, tkt := range gw.GetAssistant().GetNewPurchased() {
		-->		ticketGW := ticket.NewTicketGateway(gw.ctx, tkt)
		-->		if err = ticketGW.Update(); err != nil {
					return
				}
			}*/
	//org := Controller.GetOrganizer(clt)

	org := gw.GetOrganizer()

	//eventGW := event.NewEventGateway(gw.ctx, org.GetEventOrg)

	db.Model(org).Association("Organize").Append(org.GetEventOrg)

	return nil
}

func (gw *clientGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	return db.Delete(gw.Controller).Error
}

func (gw *clientGateway) AddFavorit() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}
	ctrl := gw.Controller.GetFavs()
	err = db.Model(gw.Controller).Association("Favs").Append(ctrl)
	return err
}
