package event

import (
	"context"

	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type txListFavorites struct {
	clientCtrl client.Controller
}

func (tx *txListFavorites) Precondition() error {
	return nil
}

// Postcondition returns the list of favorites for the user id
func (tx *txListEvents) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a List Favorites request for client %d", tx.request.ID)

	//Get the client and make sure it exists
	if tx.clientCtrl, err = clientMOD.FindClientByID(ctx, tx.request.ID) {
		return
	}
	return events, err
}

// Commit commits the transaction result
func (tx *txListFavorites) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListFavorites) Rollback() {

}
