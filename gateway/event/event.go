package event

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"gorm.io/gorm"
)

type eventGateway struct {
	*event.Event
	ctx context.Context
}

func (gw *eventGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = OpenEventStream(); err != nil {
		return
	}

	db.Create(gw.Event)
	return
}

func (gw *eventGateway) Update() (err error) {
	return nil
}

func (gw *eventGateway) Remove() (err error) {
	return nil
}
