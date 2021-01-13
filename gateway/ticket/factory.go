package ticket

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// NewTicketGateway builds a gateway for the provided client
func NewTicketGateway(ctx context.Context, ticket ticket.Controller) Gateway {
	return &ticketGateway{Controller: ticket, ctx: ctx}
}
