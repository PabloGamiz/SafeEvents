package session

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

// Controller represents a client and it's main data
type Controller interface {
	client.Controller
	context.Context
	Cookie() string
	Client() client.Controller
}
