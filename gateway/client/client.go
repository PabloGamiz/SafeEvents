package client

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type clientGateway struct {
	client.Controller
	ctx context.Context
}

func (client *clientGateway) initMongoClient() (*mongodb.Client, error) {
	return mongo.NewMongoClient(client.ctx)
}

func (client *clientGateway) Insert() (err error) {
	var c *mongodb.Client
	if c, err = client.initMongoClient(); err != nil {
		return
	}

	defer c.Disconnect(client.ctx)
	col := c.Database(mongo.Database).Collection(collection)
	var result *mongodb.InsertOneResult
	if result, err = col.InsertOne(client.ctx, client); err != nil {
		return
	}

	parsed, ok := result.InsertedID.(string)
	if !ok {
		err = fmt.Errorf(errInsertOneResultParse, result)
		return
	}

	client.SetID(parsed)
	return
}

func (client *clientGateway) Update() (err error) {
	var c *mongodb.Client
	if c, err = client.initMongoClient(); err != nil {
		return
	}

	defer c.Disconnect(client.ctx)
	return nil
}

func (client *clientGateway) Remove() (err error) {
	var c *mongodb.Client
	if c, err = client.initMongoClient(); err != nil {
		return
	}

	defer c.Disconnect(client.ctx)
	return nil
}
