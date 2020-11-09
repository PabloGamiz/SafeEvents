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
	err:= db.C(COLLECTION).InsertOne(&amp; event)
	return err
}

func (event *eventGateway) Update() (err error) {

}

func (event *eventGateway) Remove(event Event) (err error) {
	result, err := db.C(COLLECTION).DeleteOne(context.Background(), bson.D{ {"name", event.name},},)
	return err
}

func (event *eventGateway) FindAll() (err, error) {
	var c []Event
	err := c.Database(mongo.Database).Collection(collection).Find(bson.E{}).All()
	return c, err
}
