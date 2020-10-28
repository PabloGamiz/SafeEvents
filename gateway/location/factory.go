package location

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewLocationGateway builds a gateway for the provided location
func NewLocationGateway(ctx context.Context, location location.Controller) Gateway {
	return &locationGateway{Controller: location, ctx: ctx}
}

// FindLocationByID returns the gateway for the location that match the provided Id
func FindLocationByID(ctx context.Context, id string) (gw Gateway, err error) {
	var mongoClient *mongodb.Client
	if mongoClient, err = mongo.NewMongoClient(ctx); err != nil {
		return
	}

	defer mongoClient.Disconnect(ctx)
	col := mongoClient.Database(mongo.Database).Collection(collection)
	model := &location.Location{}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	if err = col.FindOne(ctx, bson.M{"_id": objectID}).Decode(model); err != nil {
		return
	}

	gw = &locationGateway{Controller: model, ctx: ctx}
	return
}
