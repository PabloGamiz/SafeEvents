package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type serviceGateway struct {
	service.Controller
	ctx   context.Context
	mongo *mongodb.Client
	mu    sync.Mutex
}

func (service *serviceGateway) iniMongoClient() (err error) {
	service.mu.Lock()
	defer service.mu.Unlock()

	if service.mongo == nil {
		var mongoClient *mongodb.Client
		if mongoClient, err = mongo.NewMongoClient(service.ctx); err == nil {
			service.mongo = mongoClient
		}
	}

	return
}

func (service *serviceGateway) getMongoClient() (mongo *mongodb.Client, err error) {
	if service.mongo == nil {
		err = service.iniMongoClient()
	}

	mongo = service.mongo
	return
}

func (service *serviceGateway) Insert() (err error) {
	var c *mongodb.Client
	if c, err = service.getMongoClient(); err != nil {
		return
	}

	col := c.Database(mongo.Database).Collection(collection)
	var result *mongodb.InsertOneResult
	if result, err = col.InsertOne(service.ctx, service); err != nil {
		return
	}

	parsed, ok := result.InsertedID.(string)
	if !ok {
		err = fmt.Errorf("Got an error while parsing InsertOneResult: %+v", result)
		return
	}
	service.setID(parsed)
	return
}

func (client *clientGateway) Update() error {
	return nil
}

func (client *clientGateway) Remove() error {
	return nil
}
