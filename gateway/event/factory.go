package event

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewEventGateway builds a gateway for the provided event
func NewEventGateway(ctx context.Context, event event.Controller) Gateway {
	return &eventGateway{Controller: event, ctx: ctx}
}

// FindEventByID returns the gateway for the event that match the provided Id
func FindEventByID(ctx context.Context, id string) (gw Gateway, err error) {
	var mongoClient *mongodb.Client
	if mongoClient, err = mongo.NewMongoClient(ctx); err != nil {
		return
	}

	defer mongoClient.Disconnect(ctx)
	col := mongoClient.Database(mongo.Database).Collection(collection)

	var model event.Event
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	if err = col.FindOne(ctx, bson.M{"_id": objectID}).Decode(model); err != nil {
		return
	}

	gw = &locationGateway{Controller: &model, ctx: ctx}
	return
}
