package ticket

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	"gorm.io/gorm"
)

type ticketGateway struct {
	ticket.Controller
	ctx context.Context
}

func (gw *ticketGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = ticket.OpenTicketStream(); err != nil {
		return
	}

	db.Create(gw.Controller)
	return
}

func (gw *ticketGateway) Update() (err error) {
	var db *gorm.DB
	if db, err = ticket.OpenTicketStream(); err != nil {
		return
	}

	db.Table("tickets").Updates(gw.Controller)
	return
}

func (gw *ticketGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = ticket.OpenTicketStream(); err != nil {
		return
	}

	db.Table("tickets").Delete(gw.Controller)
	return
}
