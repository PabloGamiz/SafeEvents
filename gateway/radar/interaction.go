package radar

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type interactionGateway struct {
	interaction.Controller
	ctx context.Context
}

func (gw *interactionGateway) Insert() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Create(gw.Controller)
	return
}

func (gw *interactionGateway) Update() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("interactions").Updates(gw.Controller)
	return
}

func (gw *interactionGateway) Remove() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("interactions").Delete(gw.Controller)
	return
}
