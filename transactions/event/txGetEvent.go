package event

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// txSignup represents an
type txGetEvent struct {
}

// Precondition validates the transaction is ready to run
func (tx *txGetEvent) Precondition(context.Context) error {
	//CHECK IF EXISTS
	var gw EventGW.Gateway
	if gw, err := gw.FindEventByName(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return err
		}
	}

}

// Postcondition creates new user and a opens its first session
func (tx *txGetEvent) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Event request for event %s", tx.info.Name)
	var gw EventGW.Gateway
	if gw, response = eventGW.FindEventByName(ctx, tx.info.Name); response == nil {
		return
	}
	return response, nil
}

// Commit commits the transaction result
func (tx *txGetEvent) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txGetEvent) Rollback() {

}
