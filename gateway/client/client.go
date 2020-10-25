package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type clientGateway struct {
	client.Controller
	ctx   context.Context
	mongo *mongodb.Client
	mu    sync.Mutex
}

func (client *clientGateway) iniMongoClient() (err error) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.mongo == nil {
		var mongoClient *mongodb.Client
		if mongoClient, err = mongo.NewMongoClient(client.ctx); err == nil {
			client.mongo = mongoClient
		}
	}

	return
}

func (client *clientGateway) getMongoClient() (mongo *mongodb.Client, err error) {
	if client.mongo == nil {
		err = client.iniMongoClient()
	}

	mongo = client.mongo
	return
}

func (client *clientGateway) Insert() (err error) {
	var c *mongodb.Client
	if c, err = client.getMongoClient(); err != nil {
		return
	}

	col := c.Database(mongo.Database).Collection(collection)
	var result *mongodb.InsertOneResult
	if result, err = col.InsertOne(client.ctx, client); err != nil {
		return
	}

	parsed, ok := result.InsertedID.(string)
	if !ok {
		err = fmt.Errorf("Got an error while parsing InsertOneResult: %+v", result)
		return
	}
	client.SetID(parsed)
	return
}

func (client *clientGateway) Update() error {
	return nil
}

func (client *clientGateway) Remove() error {
	return nil
}
