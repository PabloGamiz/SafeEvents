package mongo

import (
	"context"
	"log"
	"testing"
)

const (
	defaultMongoURL = "mongodb://mongo:8080"
)

func TestMongoClientConnection(t *testing.T) {
	if len(mongoURL) == 0 {
		log.Printf("No URI provided for mongodb, by default it is %s\n", defaultMongoURL)
		mongoURL = defaultMongoURL
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	client, err := NewMongoClient(ctx)
	if err != nil {
		t.Fatalf("Got %s; while getting the Mongo client", err.Error())
	}

	if err = client.Ping(ctx, nil); err != nil {
		t.Fatalf("Got %s; while sending ping to Mongo database", err.Error())
	}

	if err = client.Disconnect(ctx); err != nil {
		t.Fatalf("Got %s; while disconnecting from the MongoDB", err.Error())
	}
}
