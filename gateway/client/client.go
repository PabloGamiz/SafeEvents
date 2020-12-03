package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type clientGateway struct {
	client.Controller
	ctx context.Context
}

func (gw *clientGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	if err = db.Table("clients").Create(gw.Controller).Error; err != nil {
		return
	}

	assist := gw.Controller.GetAssistant()
	assist.SetParent(gw.Controller) // once the client have been inserted, it has an ID, so we must ensure the assistant keeps that id as FK
	if err = assistant.NewAssistantGateway(gw.ctx, assist).Insert(); err != nil {
		return
	}

	organ := gw.Controller.GetOrganizer()
	organ.SetParent(gw.Controller) // once the client have been inserted, it has an ID, so we must ensure the organizer keeps that id as FK
	return organizer.NewOrganizerGateway(gw.ctx, organ).Insert()
}

func (gw *clientGateway) Update() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	if db = db.Table("clients").Updates(gw.Controller); db.Error != nil {
		return db.Error
	}

	assist := gw.Controller.GetAssistant()
	if err = assistant.NewAssistantGateway(gw.ctx, assist).Update(); err != nil {
		return
	}

	organ := gw.Controller.GetOrganizer()
	return organizer.NewOrganizerGateway(gw.ctx, organ).Update()
}

func (gw *clientGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	if db = db.Table("clients").Delete(gw.Controller); db.Error != nil {
		return db.Error
	}

	assist := gw.Controller.GetAssistant()
	if err = assistant.NewAssistantGateway(gw.ctx, assist).Remove(); err != nil {
		return
	}

	organ := gw.Controller.GetOrganizer()
	return organizer.NewOrganizerGateway(gw.ctx, organ).Remove()
}

func (gw *clientGateway) AddFavorit() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}
	ctrl := gw.Controller.GetFavs()
	err = db.Model(gw.Controller).Association("Favs").Append(ctrl)
	return err
}

func (gw *clientGateway) DeleteFavorit(ctrl event.Controller) (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}
	err = db.Model(gw.Controller).Association("Favs").Delete(gw.Controller, ctrl)
	return err
}

func (gw *clientGateway) FindFavorit(ctrl event.Controller) (faved bool, err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}
	var eventsMOD []*event.Event
	err = db.Model(gw.Controller).Association("Favs").Find(&eventsMOD)
	for _, evnt := range eventsMOD {
		if evnt.GetID() == ctrl.GetID() {
			return true, err
		}
	}

	return false, err
}
