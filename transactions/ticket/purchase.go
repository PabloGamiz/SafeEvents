package ticket

import (
	"context"
	"fmt"
	"log"

	"github.com/PabloGamiz/SafeEvents-Backend/model/session"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	ticketGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txPurchase represents an
type txPurchase struct {
	request   ticketDTO.PurchaseRequestDTO
	sessCtrl  session.Controller
	purchased []ticketGW.Gateway
}

func (tx *txPurchase) buildPurchaseResponseDTO() *ticketDTO.PurchaseResponseDTO {
	ticketsID := make([]uint, len(tx.purchased))
	for index, ticket := range tx.purchased {
		ticketsID[index] = ticket.GetID()
	}

	return &ticketDTO.PurchaseResponseDTO{
		TicketsID: ticketsID,
	}
}

func (tx *txPurchase) buildNewTicket(ctx context.Context) (gw ticketGW.Gateway, err error) {
	tick := &ticketMOD.Ticket{
		ClientID: tx.sessCtrl.GetID(),
		EventID:  tx.request.EventID,
	}

	gw = ticketGW.NewTicketGateway(ctx, tick)
	if err = gw.Insert(); err != nil {
		return
	}

	if got := ticketMOD.Option(tx.request.Option); got == ticketMOD.BOUGHT {
		err = gw.Buy()
	}

	return
}

// Precondition validates the transaction is ready to run
func (tx *txPurchase) Precondition() (err error) {
	// make sure the session exists
	if tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie); err != nil {
		return
	}
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txPurchase) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Purchase request from client %v", tx.sessCtrl.GetID())

	// make sure the event exists
	var event eventGW.Gateway
	if event, err = eventGW.FindEventByID(ctx, int(tx.request.EventID)); err != nil {
		return
	}

	var tickets []ticketMOD.Controller
	if tickets, err = ticketGW.GetTicketsByEventID(tx.request.EventID); err != nil {
		return
	}

	if len(tickets)+tx.request.HowMany > event.GetCapacity() {
		err = fmt.Errorf(errNotStock)
		return
	}

	// foreach ticket to purchase
	tx.purchased = make([]ticketGW.Gateway, tx.request.HowMany)
	for it := 0; it < tx.request.HowMany; it++ {
		if tx.purchased[it], err = tx.buildNewTicket(ctx); err != nil {
			return
		}
	}

	response := tx.buildPurchaseResponseDTO()
	return response, nil
}

// Commit commits the transaction result
func (tx *txPurchase) Commit() (err error) {
	return
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPurchase) Rollback() {
	for _, ticket := range tx.purchased {
		ticket.Remove()
	}
}
