package ticket

import (
	"context"
	"fmt"
	"log"

	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	ticketGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txPurchase represents an
type txPurchase struct {
	request   ticketDTO.PurchaseRequestDTO
	sessCtrl  session.Controller
	purchased []ticketGW.Gateway
	ctx       context.Context
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
		AssistantID: tx.sessCtrl.GetID(),
		EventID:     tx.request.EventID,
	}

	gw = ticketGW.NewTicketGateway(ctx, tick)
	if got := ticketMOD.Option(tx.request.Option); got == ticketMOD.BOUGHT {
		if err = gw.Activate(); err != nil {
			return
		}
	}

	err = gw.Insert()
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

	var tickets []ticket.Controller
	if tickets, err = ticket.GetTicketsByEventID(tx.request.EventID); err != nil {
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

	tx.ctx = ctx
	response := tx.buildPurchaseResponseDTO()
	return response, nil
}

// Commit commits the transaction result
func (tx *txPurchase) Commit() (err error) {
	for _, ticket := range tx.purchased {
		tx.sessCtrl.GetAssistant().AddPurchase(ticket)
	}

	clientGW := client.NewClientGateway(tx.ctx, tx.sessCtrl)
	return clientGW.Update()
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPurchase) Rollback() {
	for _, ticket := range tx.purchased {
		ticket.Remove()
	}
}
