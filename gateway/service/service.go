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
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Create(gw.Controller)
	return
}

func (gw *serviceGateway) Update() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Save(gw.Controller)
	return
}

func (gw *serviceGateway) Remove() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Delete(gw.Controller)
	return
}
