package ticket

import (
	"context"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txActivate represents an
type txActivate struct {
	request  ticketDTO.ActivateRequestDTO
	sessCtrl session.Controller
}

// Precondition validates the transaction is ready to run
func (tx *txActivate) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txActivate) Postcondition(ctx context.Context) (v interface{}, err error) {
	return
}

// Commit commits the transaction result
func (tx *txActivate) Commit() (err error) {
	return
}

// Rollback rollbacks any change caused while the transaction
func (tx *txActivate) Rollback() {

}
