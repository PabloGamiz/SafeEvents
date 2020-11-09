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

func (tx *txPublicaEvent) Precondition(ctx context.Context) (err error) { //Comprova que no existeixi l'event
	if gw, err := eventGW.FindEventByName(ctx, bson.D{{"name", tx.requested.name}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err != mongo.ErrNoDocuments {
			return err
		}
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