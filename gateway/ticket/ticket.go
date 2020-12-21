package ticket

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type ticketGateway struct {
	ticket.Controller
	ctx context.Context
}

func (gw *ticketGateway) Insert() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Create(gw.Controller)
	return
}

func (gw *ticketGateway) Update() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("tickets").Updates(gw.Controller)
	return
}

func (gw *ticketGateway) Remove() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("tickets").Delete(gw.Controller)
	return
}
