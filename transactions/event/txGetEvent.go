package event

import (
	"context"
	"log"
)

// txSignup represents an
type txGetEvent struct {
}

// Precondition validates the transaction is ready to run
func (tx *txGetEvent) Precondition(context.Context) error {
	//CHECK IF EXISTS
	count, err := collection.EstimatedDocumentCount(context.Background)
	return
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
