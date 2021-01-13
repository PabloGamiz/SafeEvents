package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

// NewClientGateway builds a gateway for the provided client
func NewClientGateway(ctx context.Context, client client.Controller) Gateway {
	return &clientGateway{Controller: client, ctx: ctx}
}
