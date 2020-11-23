package event

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// NewEventGateway builds a gateway for the provided event
func NewEventGateway(ctx context.Context, event event.Controller) Gateway {
	return &eventGateway{Controller: event, ctx: ctx}
}
