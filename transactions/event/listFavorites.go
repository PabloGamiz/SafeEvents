package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

type txListFavorites struct {
	request    eventDTO.ListFavoritesRequestDTO
	clientCtrl client.Controller
	sessCtrl   session.Controller
}

func (tx *txListFavorites) Precondition() error {
	/*
		tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
		return
	*/
	return nil
}

// Postcondition returns the list of favorites for the user id
func (tx *txListFavorites) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a List Favorites request for client %s", tx.request.Cookie)

	// SESSION //

	var sess sessionMOD.Controller
	if sess, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		log.Printf("No id found for cookie %s", tx.request.Cookie)
		return
	}
	//Get the client and make sure it exists
	tx.clientCtrl = sess.Client()

	events, err := clientMOD.FindAllFavs(ctx, tx.clientCtrl)

	return events, err
}

// Commit commits the transaction result
func (tx *txListFavorites) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListFavorites) Rollback() {

}
