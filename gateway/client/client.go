package client

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type clientGateway struct {
	client.Controller
	ctx context.Context
}

func (client *clientGateway) Insert() error {
	return nil
}

func (client *clientGateway) Update() error {
	return nil
}

func (client *clientGateway) Remove() error {
	return nil
}
