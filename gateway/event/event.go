package event

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type eventGateway struct {
	event.Controller
	ctx context.Context
}

func (gw *eventGateway) Insert() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	return db.Create(gw.Controller).Error
}

func (gw *eventGateway) Update() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	return db.Save(gw.Controller).Error
}

func (gw *eventGateway) Remove() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	return db.Delete(gw.Controller).Error
}

func (gw *eventGateway) FindAll() (err error) {
	return nil
}
