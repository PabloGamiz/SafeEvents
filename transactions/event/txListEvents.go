package event

import (
	"context"
	"log"
)

// txSignup represents an
type txListEvents struct {
}

// Precondition validates the transaction is ready to run
func (tx *txListEvents) Precondition() error {
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txListEvents) Postcondition(context.Context) (interface{}, error) {
	log.Printf("Got a List Events request")
	response := "test" // TODO: Set a fake EventDTO as the response and test the api call on postman.

	return response, nil
}

// Commit commits the transaction result
func (tx *txListEvents) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListEvents) Rollback() {

}
