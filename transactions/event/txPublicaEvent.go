package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request  eventDTO.DTO
	sessCtrl session.Controller
}

func (tx *txPublicaEvent) Precondition() (err error) { //Comprova que no existeixi l'event
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

func (tx *txPublicaEvent) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a Publica Event request for event %s", tx.request.Title)

	event := &eventMOD.Event{
		Title:       tx.request.Title,
		Description: tx.request.Description,
		Capacity:    tx.request.Capacity,
		Price:       tx.request.Price,
		CheckInDate: tx.request.CheckInDate,
		ClosureDate: tx.request.ClosureDate,
		Location:    tx.request.Location,
	}

	gw := eventGW.NewEventGateway(ctx, evnt)
	err := gw.Insert()
	log.Println(gw.GetID())
	return gw, err
}

func (tx *txPublicaEvent) Commit() error {
	return nil
}

func (tx *txPublicaEvent) Rollback() {

}
