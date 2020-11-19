package client

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
)

// Controller represents a client and it's main data
type Controller interface {
	GetID() uint
	GetEmail() string
	GetAssistant() assistant.Controller
	GetOrganizer() organizer.Controller
}
