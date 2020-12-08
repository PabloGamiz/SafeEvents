package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// txListEvents represents the list events transaction
type txListEventsByType struct {
	request eventDTO.ListEventsByTypeRequestDTO
}

// Precondition validates the transaction is ready to run
func (tx *txListEventsByType) Precondition() error {
	return nil
}

// Postcondition lists events from the database
func (tx *txListEventsByType) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a List Events request")
	events, err := eventMOD.FindAllByType(ctx, tx.request.EventType)
	return events, err
}

// Commit commits the transaction result
func (tx *txListEventsByType) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListEventsByType) Rollback() {

}
