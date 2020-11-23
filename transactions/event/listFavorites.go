package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type txListFavorites struct {
	request    eventDTO.ListFavoritesRequestDTO
	clientCtrl client.Controller
}

func (tx *txListFavorites) Precondition() error {
	return nil
}

// Postcondition returns the list of favorites for the user id
func (tx *txListFavorites) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a List Favorites request for client %d", tx.request.ID)

	//Get the client and make sure it exists
	if tx.clientCtrl, err = clientMOD.FindClientByID(ctx, tx.request.ID); err != nil {
		return
	}

	events := tx.clientCtrl.GetFavs()

	return events, err
}

// Commit commits the transaction result
func (tx *txListFavorites) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListFavorites) Rollback() {

}
