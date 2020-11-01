package event

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewEventGateway builds a gateway for the provided event
func NewEventGateway(ctx context.Context, event event.Controller) Gateway {
	return &NewEventGateway{Controller: event, ctx: ctx}
}

// FindEventByName returns the gateway for the event that match the provided name
func FindEventByName(ctx context.Context, name string) (gw Gateway, err error) {
	var mongoEvent *mongodb.Event
	if mongoEvent, err = mongo.NewMongoEvent(ctx); err != nil {
		return
	}

	defer mongoEvent.Disconnect(ctx)
	col := mongoEvent.Database(mongo.Database).Collection(collection)

	var model event.Event
	if err = col.FindOne(ctx, bson.M{"name": name}).Decode(&model); err != nil {
		err = fmt.Errorf("Got error %s, while searching for name %s", err.Error(), email)
		return
	}

	gw = &eventGateway{Controller: &model, ctx: ctx}
	return
}
