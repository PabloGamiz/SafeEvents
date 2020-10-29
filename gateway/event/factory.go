package event

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewEventGateway builds a gateway for the provided event
func NewEventGateway(ctx context.Context, event event.Controller) Gateway {
	return &NewEventGateway{Controller: event, ctx: ctx}
}

// FindEventByName ???
func FindEventByID(ctx context.Context, id string) (gw Gateway, err error) {
	
	col .= mongoEvent.Database(mongo.Database).Collection(collection)
	model :=  &event.Event{}
	if err = col.FindOne(ctx, bson.M{ /*"id": id*/ }).Decode(model); err != nil {
		return
	}

	//gw = &eventGateway{Controller:model, ctx:ctx}
	//return
}