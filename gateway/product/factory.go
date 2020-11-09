package product

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// NewProductGateway builds a gateway for the provided product
func NewProductGateway(ctx context.Context, product product.Controller) Gateway {
	return &productGateway{Controller: product, ctx: ctx}
}

// FindProductByID returns the gateway for the product that match the provided Id
func FindProductByID(ctx context.Context, id string) (gw Gateway, err error) {
	var mongoClient *mongodb.Client
	if mongoClient, err = mongo.NewMongoClient(ctx); err != nil {
		return
	}

	defer mongoClient.Disconnect(ctx)
	col := mongoClient.Database(mongo.Database).Collection(collection)

	var model product.Product
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	if err = col.FindOne(ctx, bson.M{"_id": objectID}).Decode(model); err != nil {
		return
	}

	gw = &productGateway{Controller: &model, ctx: ctx}
	return
}
