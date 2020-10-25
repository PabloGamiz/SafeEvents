package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURL = os.Getenv(EnvMongoURL)
)

// NewMongoClient returns a brand new client
func NewMongoClient(ctx context.Context) (client *mongo.Client, err error) {
	mongoCtx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	options := options.Client().ApplyURI(mongoURL)
	client, err = mongo.Connect(mongoCtx, options)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewDatabaseConnection returns a brand new database connection
func NewDatabaseConnection(ctx context.Context) (db *mongo.Database, err error) {
	var client *mongo.Client
	if client, err = NewMongoClient(ctx); err != nil {
		return
	}

	db = client.Database(Database)
	return
}
