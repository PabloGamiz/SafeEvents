package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
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

	for _, tkt := range gw.GetAssistant().GetNewPurchased() {
		ticketGW := ticket.NewTicketGateway(gw.ctx, tkt)
		if err = ticketGW.Update(); err != nil {
			return
		}
	}

	return nil
}

func (gw *clientGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	return db.Delete(gw.Controller).Error
}
