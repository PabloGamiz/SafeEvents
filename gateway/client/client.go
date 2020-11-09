package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"gorm.io/gorm"
)

type clientGateway struct {
	client.Controller
	ctx context.Context
}

func (gw *clientGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = OpenClientStream(); err != nil {
		return
	}

	db.Create(gw.Controller)
	return
}

func (gw *clientGateway) Update() (err error) {
	return nil
}

func (gw *clientGateway) Remove() (err error) {
	return nil
}
