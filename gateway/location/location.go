package location

import (
	"context"
	"fmt"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type locationGateway struct {
	location.Controller
	ctx   context.Context
	mongo *mongodb.Client
	mu    sync.Mutex
}

func (location *locationGateway) iniMongoClient() (err error) {
	location.mu.Lock()
	defer location.mu.Unlock()

	if location.mongo == nil {
		var mongoClient *mongodb.Client
		if mongoClient, err = mongo.NewMongoClient(location.ctx); err == nil {
			location.mongo = mongoClient
		}
	}

	return
}

func (location *locationGateway) getMongoClient() (mongo *mongodb.Client, err error) {
	if location.mongo == nil {
		err = location.iniMongoClient()
	}

	mongo = location.mongo
	return
}

func (location *locationGateway) Insert() (err error) {
	var c *mongodb.Client
	if c, err = location.getMongoClient(); err != nil {
		return
	}

	col := c.Database(mongo.Database).Collection(collection)
	var result *mongodb.InsertOneResult
	if result, err = col.InsertOne(location.ctx, location); err != nil {
		return
	}

	parsed, ok := result.InsertedID.(string)
	if !ok {
		err = fmt.Errorf("Got an error while parsing InsertOneResult: %+v", result)
		return
	}
	location.SetID(parsed)
	return
}

func (location *locationGateway) Update() error {
	return nil
}

func (location *locationGateway) Remove() error {
	return nil
}
