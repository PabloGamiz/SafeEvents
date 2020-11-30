package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request    eventDTO.DTO
	sessCtrl   session.Controller
	clientCtrl client.Controller
	ctx        context.Context
	eventCtrl  event.Controller
}

func (tx *txPublicaEvent) Precondition() (err error) { //Comprova que no existeix l'event
	// make sure the session exists
	//tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

func (tx *txPublicaEvent) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Publica Event request for event %s", tx.request.Title)

	eventCtrl := &eventMOD.Event{
		Title:       tx.request.Title,
		Description: tx.request.Description,
		Capacity:    tx.request.Capacity,
		Price:       tx.request.Price,
		CheckInDate: tx.request.CheckInDate,
		Price:       tx.request.Price,
		ClosureDate: tx.request.ClosureDate,
		Location:    tx.request.Location,
		Image:       tx.request.Image,
		Tipus:       tx.request.Tipus,
	}
	gw := eventGW.NewEventGateway(ctx, eventCtrl)
	err = gw.Insert()
	log.Println(gw.GetID())
	tx.ctx = ctx
	log.Println(err)

	var ctr client.Controller
	ctr, err = clientMOD.FindClientByID(tx.ctx, 2)
	ctr.GetOrganizer().AddEvent(eventCtrl)
	clientgw := clientGW.NewClientGateway(tx.ctx, ctr)
	if err = clientgw.Update(); err != nil {
		return
	}
	return gw, err
}

func (tx *txPublicaEvent) Commit() (err error) {

	return
}

func (tx *txPublicaEvent) Rollback() {
	clientgw := clientGW.NewClientGateway(tx.ctx, tx.sessCtrl)
	clientgw.Remove()
}
