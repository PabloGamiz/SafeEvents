package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// txSignup represents an
type txGetEvent struct {
	request eventDTO.DTO
}

// Precondition validates the transaction is ready to run
func (tx *txGetEvent) Precondition() error {
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txGetEvent) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a Event request for event with and ID of %d ", tx.request.ID)
	event, err := eventMOD.FindEventByID(ctx, tx.request.ID)
	return event, err
}

// Commit commits the transaction result
func (tx *txGetEvent) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txGetEvent) Rollback() {

}
