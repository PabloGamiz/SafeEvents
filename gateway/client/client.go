package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
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

	org := gw.GetOrganizer()

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

func (gw *clientGateway) DeleteFavorit(ctrl event.Controller) (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}
	err = db.Model(gw.Controller).Association("Favs").Delete(gw.Controller, ctrl)
	return err
}

func (gw *clientGateway) FindFavorit(ctrl event.Controller) (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}
	err = db.Model(gw.Controller).Association("Favs").Find(gw.Controller, ctrl)
	return err
}
