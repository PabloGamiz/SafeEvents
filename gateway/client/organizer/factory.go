package organizer

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
)

// NewOrganizerGateway builds a gateway for the provided client
func NewOrganizerGateway(ctx context.Context, organizer organizer.Controller) Gateway {
	return &organizerGateway{Controller: organizer, ctx: ctx}
}
