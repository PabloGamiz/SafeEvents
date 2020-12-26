package session

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	radarMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar"
)

// Controller represents a client and it's main data
type Controller interface {
	client.Controller
	context.Context
	Cookie() string
	Client() client.Controller
	GetRadar() radarMOD.Controller
	InitRadar(string) error
	FinishRadar() error
}
