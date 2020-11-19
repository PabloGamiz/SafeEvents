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

	db.Create(gw.Controller)
	return
}

func (gw *clientGateway) Update() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	db.Save(gw.Controller)
	for _, tkt := range gw.GetAssistant().GetNewPurchased() {
		ticketGW := ticket.NewTicketGateway(gw.ctx, tkt)
		ticketGW.Update()
	}

	return nil
}

func (gw *clientGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	db.Delete(gw.Controller)
	return nil
}
