package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request    eventDTO.DTO
	sessCtrl   session.Controller
	clientCtrl client.Controller
	gwClt      clientGW.Gateway
	ctx        context.Context
	eventCtrl  event.Controller
}

func (tx *txPublicaEvent) Precondition() (err error) { //Comprova que no existeixi l'event
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

func (tx *txPublicaEvent) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Publica Event request for event %s", tx.request.Title)

	eventCtrl := &eventMOD.Event{
		Title:       tx.request.Title,
		Description: tx.request.Description,
		Capacity:    tx.request.Capacity,
		CheckInDate: tx.request.CheckInDate,
		Price:       tx.request.Price,
		ClosureDate: tx.request.ClosureDate,
		Location:    tx.request.Location,
	}
	gw := eventGW.NewEventGateway(ctx, eventCtrl)
	err = gw.Insert()
	log.Println(gw.GetID())
	tx.ctx = ctx
	return gw, err
}

func (tx *txPublicaEvent) Commit() (err error) {
	if tx.clientCtrl, err = clientMOD.FindEventByID(ctx, sessCtrl)
	clientGW := clientGW.NewClientGateway(tx.ctx, clientCtrl)
	tx.sessCtrl.GetOrganizer().AddEvent(clientGW)
	if err = clientGW.Update(); err != nil {
		return
	}

	eventGW := eventGW.NewEventGateway(tx.ctx, tx.eventCtrl)
	if err = eventGW.Update(); err != nil {
		return
	}
}
func (tx *txPublicaEvent) Rollback() {
	gwClt.Remove()
}
