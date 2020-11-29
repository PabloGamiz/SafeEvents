package ticket

import (
	"context"
	"log"
	"time"

	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	ticketGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txPurchase represents an
type txPurchase struct {
	request   ticketDTO.PurchaseRequestDTO
	sessCtrl  session.Controller
	purchased []ticketGW.Gateway
	eventCtrl event.Controller
	ctx       context.Context
	taked     bool
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
	assistID := tx.sessCtrl.GetAssistant().GetID()
	tick := &ticketMOD.Ticket{
		AssistantID: assistID,
		EventID:     tx.request.EventID,
		CreatedAt:   time.Now(),
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

func (tx *txPurchase) releaseBookedTickets() (err error) {
	var tickets []ticket.Controller
	if tickets, err = ticketMOD.GetTicketsByEventID(tx.eventCtrl.GetID()); err != nil {
		return
	}

	released := 0
	for _, ticket := range tickets {
		if ticket.GetOption() == ticketMOD.BOOKED {
			created := ticket.GetCreatedAt()
			if diff := time.Now().Sub(created); diff > 24*time.Hour {
				gwTicket := ticketGW.NewTicketGateway(tx.ctx, ticket)
				if err = gwTicket.Remove(); err != nil {
					break // not returning, due n could be greater than 0; so the TakenTickets substract must be done
				}

				released++
			}
		}
	}

	tx.eventCtrl.TakeTickets(-released)
	return
}

// Precondition validates the transaction is ready to run
func (tx *txPurchase) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txPurchase) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Purchase request from client %v", tx.sessCtrl.GetID())

	// make sure the event exists
	if tx.eventCtrl, err = eventMOD.FindEventByID(ctx, tx.request.EventID); err != nil {
		return
	}

	// making sure there are enought tikets to purchase
	if err = tx.eventCtrl.TakeTickets(tx.request.HowMany); err != nil {
		// if there are no enought tickets, try to release the booked ones
		if err = tx.releaseBookedTickets(); err != nil {
			return
		}

		// try again to purchase tickets
		if err = tx.eventCtrl.TakeTickets(tx.request.HowMany); err != nil {
			return
		}
	}

	// foreach ticket to purchase
	tx.taked = true
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
	// adding tickets to user
	for _, ticket := range tx.purchased {
		tx.sessCtrl.GetAssistant().AddPurchase(ticket)
	}

	eventgw := eventGW.NewEventGateway(tx.ctx, tx.eventCtrl)
	if err = eventgw.Update(); err != nil {
		return
	}

	clientgw := clientGW.NewClientGateway(tx.ctx, tx.sessCtrl)
	return clientgw.Update()
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPurchase) Rollback() {
	if tx.taked {
		tx.eventCtrl.TakeTickets(-tx.request.HowMany)
	}

	for _, ticket := range tx.purchased {
		ticket.Remove()
	}
}
