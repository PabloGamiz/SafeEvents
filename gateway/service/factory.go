package service

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewServiceGateway builds a gateway for the provided service
func NewServiceGateway(ctx context.Context, service service.Controller) Gateway {
	return &serviceGateway{Controller: service, ctx: ctx}
}

// FindServiceByID returns the gateway for the product that match the provided Id
func FindServiceByID(ctx context.Context, id primitive.ObjectID) (gw Gateway, err error) {
	var mongoClient *mongodb.Client
	if mongoClient, err = mongo.NewMongoClient(ctx); err != nil {
		return
	}

	defer mongoClient.Disconnect(ctx)
	col := mongoClient.Database(mongo.Database).Collection(collection)

	var model service.Service
	if err = col.FindOne(ctx, bson.M{"_id": id}).Decode(model); err != nil {
		return
	}

	gw = &serviceGateway{Controller: &model, ctx: ctx}
	return
}
