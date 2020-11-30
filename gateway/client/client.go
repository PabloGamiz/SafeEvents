package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client/organizer"
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

	if err = db.Create(gw.Controller).Error; err != nil {
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
	if db, err = client.OpenClientStream(); err != nil {
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
	if db, err = client.OpenClientStream(); err != nil {
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
