package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request eventDTO.DTO
}

func (tx *txPublicaEvent) publicaNewEvent(ctx context.Context) (err error) {
	gw := eventGW.NewEventGateway(ctx, request)
	return gw.Insert()
}

func (tx *txPublicaEvent) Precondition(ctx context.Context) (err error) { //Comprova que no existeixi l'event
	return
}

func (tx *txPublicaEvent) PostCondition(ctx context.Context) (err error) {
	log.Printf("Got a Publica Event request for event %s", tx.request.Title)

	var gw eventGW.Gateway
	if gw, err = tx.publicaNewEvent(ctx); err != nil {
		return
	}

	response := &eventDTO.DTO
	return response, nil
}

func (tx *txPublicaEvent) commit() error {
	return nil
}

func (tx *txPublicaEvent) Rollback() {

}
