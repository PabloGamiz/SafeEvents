package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewClientGateway builds a gateway for the provided client
func NewClientGateway(ctx context.Context, client client.Controller) Gateway {
	return &clientGateway{Controller: client, ctx: ctx}
}

// FindClientByEmail returns the gateway for the client that match the provided mail
func FindClientByEmail(ctx context.Context, email string) (gw Gateway, err error) {
	var mongoClient *mongodb.Client
	if mongoClient, err = mongo.NewMongoClient(ctx); err != nil {
		return
	}

	defer mongoClient.Disconnect(ctx)
	col := mongoClient.Database(mongo.Database).Collection(collection)
	model := &client.Client{}

	if err = col.FindOne(ctx, bson.M{"email": email}).Decode(model); err != nil {
		return
	}

	gw = &clientGateway{Controller: model, ctx: ctx}
	return
}
