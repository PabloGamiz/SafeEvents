package event

import (
	"context"
	"log"

	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// txListEvents represents the list events transaction
type txListEvents struct {
}

// Precondition validates the transaction is ready to run
func (tx *txListEvents) Precondition() error {
	return nil
}

// Postcondition lists events from the database
func (tx *txListEvents) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a List Events request")
	events, err := eventMOD.FindAll(ctx)
	return events, err
}

// Commit commits the transaction result
func (tx *txListEvents) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListEvents) Rollback() {

}
