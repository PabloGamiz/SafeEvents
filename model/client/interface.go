package client

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// Controller represents a client and it's main data
type Controller interface {
	GetID() uint
	GetEmail() string
	GetAssistant() assistant.Controller
	GetOrganizer() organizer.Controller
	AddFav(ctrl *event.Event)
	RemoveFav(ctrl *event.Event)
	GetFavs() []*event.Event
}
