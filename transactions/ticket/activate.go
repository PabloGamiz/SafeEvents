package ticket

import (
	"context"
	"fmt"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	ticketGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txActivate represents an
type txActivate struct {
	request    ticketDTO.ActivateRequestDTO
	sessCtrl   session.Controller
	toActivate []ticketMOD.Controller
	ctx        context.Context
}

func (tx *txActivate) buildPurchaseResponseDTO() *ticketDTO.PurchaseResponseDTO {
	tickets := make([]ticket.Controller, len(tx.toActivate))
	for index, ticket := range tx.toActivate {
		tickets[index] = ticket
	}

	return &ticketDTO.PurchaseResponseDTO{
		Tickets: tickets,
	}
}

// Precondition validates the transaction is ready to run
func (tx *txActivate) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txActivate) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got an Activation request from client %v", tx.sessCtrl.GetID())

	tickets := tx.sessCtrl.GetAssistant().GetPurchased()
	tx.toActivate = make([]ticket.Controller, tx.request.HowMany)
	found := 0

	for _, ticket := range tickets {
		if found == tx.request.HowMany {
			break
		}

		if ticket.GetEventID() == tx.request.EventID &&
			ticket.GetOption() == ticketMOD.BOOKED {
			tx.toActivate[found] = ticket
			found++
		}
	}

	if found != tx.request.HowMany {
		err = fmt.Errorf(errActivate, tx.request.HowMany, found)
		return
	}

	for _, ticket := range tx.toActivate {
		if err = ticket.Activate(); err != nil {
			return
		}
	}

	tx.ctx = ctx
	response := tx.buildPurchaseResponseDTO()
	return response, nil
}

// Commit commits the transaction result
func (tx *txActivate) Commit() (err error) {
	// updating tickets to database
	for _, ticket := range tx.toActivate {
		gate := ticketGW.NewTicketGateway(tx.ctx, ticket)
		if err = gate.Update(); err != nil {
			return
		}
	}

	return
}

// Rollback rollbacks any change caused while the transaction
func (tx *txActivate) Rollback() {

}
