package organizer

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"gorm.io/gorm"
)

type organizerGateway struct {
	organizer.Controller
	ctx context.Context
}

func (gw *organizerGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	return db.Create(gw.Controller).Error
}

func (gw *organizerGateway) Update() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	return db.Table("organizers").Updates(gw.Controller).Error
}

func (gw *organizerGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = client.OpenClientStream(); err != nil {
		return
	}

	return db.Table("organizers").Delete(gw.Controller).Error
}
