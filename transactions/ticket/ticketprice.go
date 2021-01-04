package ticket

import (
	"context"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txTicketPrice represents an
type txTicketPrice struct {
	request    ticketDTO.PriceTicketRequestDTO
	sessCtrl   session.Controller
	ticketCtrl ticket.Controller
	ctx        context.Context
}

func (tx *txTicketPrice) buildPriceTicketResponseDTO() *ticketDTO.PriceTicketResponseDTO {
	var evtCtrl event.Controller
	evtCtrl, _ = eventMOD.FindEventByID(tx.ticketCtrl.GetEventID())
	price := evtCtrl.GetPrice()
	return &ticketDTO.PriceTicketResponseDTO{
		Price: price,
	}
}

// Precondition validates the transaction is ready to run
func (tx *txTicketPrice) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txTicketPrice) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Paypal Purchase request from client %v", tx.sessCtrl.GetID())

	//make sure the ticket exists
	if tx.ticketCtrl, err = ticketMOD.GetTicketByID(tx.request.TicketID); err != nil {
		return
	}
	response := tx.buildPriceTicketResponseDTO()
	return response, nil
}

// Commit commits the transaction result
func (tx *txTicketPrice) Commit() (err error) {
	return err
}

// Rollback rollbacks any change caused while the transaction
func (tx *txTicketPrice) Rollback() {

}
