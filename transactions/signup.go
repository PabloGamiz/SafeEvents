package transactions

import (
	"context"
	"log"
)

// txSignup represents an
type txSignup struct {
}

// Precondition validates the transaction is ready to run
func (tx *txSignup) Precondition() error {
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txSignup) Postcondition(context.Context) (interface{}, error) {
	log.Printf("Got a Signup request")
	return nil, nil
}

// Commit commits the transaction result
func (tx *txSignup) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSignup) Rollback() {

}
