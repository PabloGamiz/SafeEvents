package assistant

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type assistantGateway struct {
	assistant.Controller
	ctx context.Context
}

func (gw *assistantGateway) Insert() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	return db.Table("assistants").Create(gw.Controller).Error
}

func (gw *assistantGateway) Update() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	if db = db.Table("assistants").Updates(gw.Controller); db.Error != nil {
		return db.Error
	}

	for _, tkt := range gw.GetPurchased() {
		ticketGW := ticket.NewTicketGateway(gw.ctx, tkt)
		if err = ticketGW.Update(); err != nil {
			return
		}
	}

	return nil
}

func (gw *assistantGateway) Remove() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	return db.Table("assistants").Delete(gw.Controller).Error
}
