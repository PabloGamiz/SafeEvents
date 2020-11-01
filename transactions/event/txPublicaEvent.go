package event

import (
	"context"
	"log"
	"time"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request eventDTO.EventRequestDTO
}

func (tx *txPublicaEvent) publicaNewEvent(ctx context.Context) (err error) {
	gw := eventGW.NewEventGateway(ctx, request)
	return gw.Insert()
}

func (tx *txPublicaEvent) Precondition() (err error) { //Comprova que existeixi l'event
	if gw, err = eventGW.FindEventByID(ctx, tx.request.ID) err != nil {
		return
	}
}

func (tx *txPublicaEvent) PostCondition() (err error) {
	log.Printf("New Event Created")
	response := "test"

	return response, nil
}

func (tx *txPublicaEvent) commit() error {
	return nil
}

func (tx *txPublicaEvent) Rollback() {

}