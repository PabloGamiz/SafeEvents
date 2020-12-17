package ticket

import (
	"context"
	"fmt"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	ticketGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
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
	log.Printf("Got a Check request from organizer %v", tx.sessCtrl.GetID())

	// Checking the oragnizer does organize the provided event id
	organizer := tx.sessCtrl.GetOrganizer()
	var targetEvent eventMOD.Controller

	// looking for the target event in all events organized by the user
	for _, v := range organizer.GetEventOrg() {
		if targetEvent != nil {
			break
		}

		targetEvent = v
		if targetEvent.GetID() != tx.request.EventID {
			targetEvent = nil
		}
	}

	if targetEvent == nil {
		// if no event was found with the provided ID it means the target event is not organized by this client
		err = fmt.Errorf(errOrganizer, tx.request.EventID)
		return
	}

	// loocking for the provided qr
	if tx.ticket, err = ticketMOD.GetTicketByQR(tx.request.Qr); err != nil {
		return
	}

	// looking for the gotten event
	if eventID := tx.ticket.GetEventID(); eventID != tx.request.EventID {
		// if the ticket event does not match with the selected event
		err = fmt.Errorf(errBelongs, tx.request.EventID, eventID)
		return
	}

	// making sure the ticket's owner exists
	clientID := tx.ticket.GetClientID()
	if tx.client, err = clientMOD.FindClientByID(ctx, clientID); err != nil {
		return
	}

	// checking the ticket
	if err = tx.ticket.Check(); err != nil {
		return
	}

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
