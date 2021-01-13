package service

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
)

// NewServiceGateway builds a gateway for the provided service
func NewServiceGateway(ctx context.Context, service service.Controller) Gateway {
	return &serviceGateway{Controller: service, ctx: ctx}
}
