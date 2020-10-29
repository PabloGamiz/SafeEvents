package location

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type locationGateway struct {
	location.Controller
	ctx context.Context
}

func (gw *locationGateway) Insert() (err error) {
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

func (gw *locationGateway) Update() (err error) {
	var c *mongodb.Client
	if c, err = mongo.NewMongoClient(gw.ctx); err != nil {
		return
	}

	defer c.Disconnect(gw.ctx)
	return nil
}

func (gw *locationGateway) Remove() (err error) {
	var c *mongodb.Client
	if c, err = mongo.NewMongoClient(gw.ctx); err != nil {
		return
	}

	defer c.Disconnect(gw.ctx)
	return nil
}
