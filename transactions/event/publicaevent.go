package event

import (
	"context"
	"log"
	"time"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	"github.com/PabloGamiz/SafeEvents-Backend/google"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request eventDTO.EventRequestDTO
}

func (tx *txPublicaEvent) publicaNewEvent(ctx context.Context) (err error) {
	evnt := &eventMOD.Event{
		Nom: tx.info.Nom,
		
	}
}

func (tx *txPublicaEvent) Precondition() (err error) {
	if gw, err = eventGW.FindEventByID(ctx, tx.request.ID) err != nil {
		return
	}
}

func (tx *txPublicaEvent) PostCondition() (err error) {
	tx.
}