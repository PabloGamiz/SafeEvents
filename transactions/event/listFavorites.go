package event

import (

)

type txListFavorites struct {
}

func (tx *txListFavorites) Precondition() error {
	return nil
}

// Postcondition returns the list of favorites for the user id
func (tx *txListEvents) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a List Favorites request for client %d", tx.request.ID)
	events, err := eventGW.
	return events, err
}

// Commit commits the transaction result
func (tx *txListFavorites) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txListFavorites) Rollback() {

}
