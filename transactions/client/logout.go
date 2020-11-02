package client

import (
	"context"
	"log"
)

// txLogout represents an
type txLogout struct {
}

// Precondition validates the transaction is ready to run
func (tx *txLogout) Precondition() (err error) {
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txLogout) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Logout request for client" /*tx.info.Email*/)
	return
}

// Commit commits the transaction result
func (tx *txLogout) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txLogout) Rollback() {

}
