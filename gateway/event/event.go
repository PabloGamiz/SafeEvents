package event

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type eventGateway struct {
	event.Controller
	ctx context.Context
}

func (event *eventGateway) Insert(event Event) (err error) {
	err:= db.C(COLLECTION).Insert(&amp; event)
	return err
}

func (event *eventGateway) Update() (err error) {

}

func (event *eventGateway) Remove() (err error) {

}

func (event *eventGateway) FindAll() (err, error) {
	var c []Event
	err := c.Database(mongo.Database).Collection(collection).Find(bson.E{}).All()
	return c, err
}
