package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoClient() (interface{}, error) {
	return NewMongoClient()
}

// NewMongoClient returns a brand new client
func NewMongoClient() (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	options := options.Client().ApplyURI(mongoURL)
	client, err = mongo.Connect(ctx, options)

	if err != nil {
		return nil, err
	}

	return client, nil
}
