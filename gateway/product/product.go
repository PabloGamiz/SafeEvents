package product

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type productGateway struct {
	product.Controller
	ctx context.Context
}

func (gw *productGateway) Insert() (err error) {
	var c *mongodb.Client
	if c, err = mongo.NewMongoClient(gw.ctx); err != nil {
		return
	}

	defer c.Disconnect(gw.ctx)
	col := c.Database(mongo.Database).Collection(collection)
	var result *mongodb.InsertOneResult
	if result, err = col.InsertOne(gw.ctx, gw.Controller); err != nil {
		return
	}

	parsed, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err = fmt.Errorf(errInsertOneResultParse, result)
		return
	}

	gw.SetID(parsed)
	return
}

func (gw *productGateway) Update() (err error) {
	var c *mongodb.Client
	if c, err = mongo.NewMongoClient(gw.ctx); err != nil {
		return
	}

	defer c.Disconnect(gw.ctx)
	return nil
}

func (gw *productGateway) Remove() (err error) {
	var c *mongodb.Client
	if c, err = mongo.NewMongoClient(gw.ctx); err != nil {
		return
	}

	defer c.Disconnect(gw.ctx)
	return nil
}
