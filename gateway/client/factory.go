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
	return &clientGateway{client, ctx}
}

// FindClientByEmail returns the gateway for the client that match the provided mail
func FindClientByEmail(ctx context.Context, email string) (gw Gateway, err error) {
	var database *mongodb.Database
	if database, err = mongo.NewDatabaseConnection(ctx); err != nil {
		return
	}

	col := database.Collection(collection)
	var model client.Client
	if err = col.FindOne(ctx, bson.M{"email": email}).Decode(&model); err != nil {
		return
	}

	gw = &clientGateway{&model, ctx}
	return
}
