package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txListEvents represents the list events transaction
type txRecomanaEvents struct {
	request  eventDTO.ListFavoritesRequestDTO
	sessCtrl session.Controller
}

// Precondition validates the transaction is ready to run
func (tx *txRecomanaEvents) Precondition() (err error) {
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition lists recomended events from the database
func (tx *txRecomanaEvents) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a Recomana Events request")
	var ctrID = tx.sessCtrl.GetID()
	events, err := eventMOD.FindRecomended(ctx, ctrID)
	return events, err
}

// Commit commits the transaction result
func (tx *txRecomanaEvents) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txRecomanaEvents) Rollback() {

}
