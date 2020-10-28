package product

import (
	"context"
	"fmt"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type productGateway struct {
	product.Controller
	ctx   context.Context
	mongo *mongodb.Client
	mu    sync.Mutex
}

func (product *productGateway) iniMongoClient() (err error) {
	product.mu.Lock()
	defer product.mu.Unlock()

	if product.mongo == nil {
		var mongoClient *mongodb.Client
		if mongoClient, err = mongo.NewMongoClient(product.ctx); err == nil {
			product.mongo = mongoClient
		}
	}

	return
}

func (product *productGateway) getMongoClient() (mongo *mongodb.Client, err error) {
	if product.mongo == nil {
		err = product.iniMongoClient()
	}

	mongo = product.mongo
	return
}

func (product *productGateway) Insert() (err error) {
	var c *mongodb.Client
	if c, err = product.getMongoClient(); err != nil {
		return
	}

	col := c.Database(mongo.Database).Collection(collection)
	var result *mongodb.InsertOneResult
	if result, err = col.InsertOne(product.ctx, product); err != nil {
		return
	}

	parsed, ok := result.InsertedID.(string)
	if !ok {
		err = fmt.Errorf("Got an error while parsing InsertOneResult: %+v", result)
		return
	}
	product.SetID(parsed)
	return
}

func (product *productGateway) Update() error {
	return nil
}

func (product *productGateway) Remove() error {
	return nil
}
