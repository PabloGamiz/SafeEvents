package ticket

import (
	"context"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txGetTickets represents an
type txGetTickets struct {
	request  ticketDTO.GetTicketsRequestDTO
	sessCtrl session.Controller
}

// Precondition validates the transaction is ready to run
func (tx *txGetTickets) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txGetTickets) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Get Tickets request from client %v", tx.sessCtrl.GetID())

	eventID := tx.request.EventID
	if v = tx.sessCtrl.GetAssistant().GetPurchased(); eventID == 0 {
		return
	}

	return ticketMOD.GetTicketsByEventIDAndClientID(eventID, tx.sessCtrl.GetID())
}

// Commit commits the transaction result
func (tx *txGetTickets) Commit() (err error) {
	return
}

// Rollback rollbacks any change caused while the transaction
func (tx *txGetTickets) Rollback() {
}
