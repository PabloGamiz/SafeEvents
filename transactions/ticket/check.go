package ticket

import (
	"context"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	ticketGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txCheck represents an
type txCheck struct {
	request  ticketDTO.CheckRequestDTO
	sessCtrl session.Controller
	ticket   ticket.Controller
	client   client.Controller
	ctx      context.Context
}

func (tx *txCheck) buildCheckResponseDTO() *ticketDTO.CheckResponseDTO {
	return &ticketDTO.CheckResponseDTO{
		Ok:     true,
		Client: tx.client,
		Ticket: tx.ticket,
	}
}

// Precondition validates the transaction is ready to run
func (tx *txCheck) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txCheck) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Check request from client %v", tx.sessCtrl.GetID())

	tx.ctx = ctx
	response := tx.buildCheckResponseDTO()
	return response, nil
}

// Commit commits the transaction result
func (tx *txCheck) Commit() (err error) {
	// updating tickets to database
	gate := ticketGW.NewTicketGateway(tx.ctx, tx.ticket)
	if err = gate.Update(); err != nil {
		return
	}

	return
}

// Rollback rollbacks any change caused while the transaction
func (tx *txCheck) Rollback() {

}
