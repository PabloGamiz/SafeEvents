package service

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type serviceGateway struct {
	service.Controller
	ctx context.Context
}

func (gw *serviceGateway) Insert() (err error) {
	var db *gorm.DB
	var disconnect mysql.Disconnect
	if db, disconnect, err = mysql.OpenStream(); err != nil {
		return
	}

	defer disconnect()
	db.Create(gw.Controller)
	return
}

func (gw *serviceGateway) Update() (err error) {
	var db *gorm.DB
	var disconnect mysql.Disconnect
	if db, disconnect, err = mysql.OpenStream(); err != nil {
		return
	}

	defer disconnect()
	db.Save(gw.Controller)
	return
}

func (gw *serviceGateway) Remove() (err error) {
	var db *gorm.DB
	var disconnect mysql.Disconnect
	if db, disconnect, err = mysql.OpenStream(); err != nil {
		return
	}

	defer disconnect()
	db.Delete(gw.Controller)
	return
}
