package client

import (
	AssistantMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	OrganizerMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	EventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// ClientInfoRequestDTO represents the expected data from a Client Info request
type ClientInfoRequestDTO struct {
	ID     uint   `json:"id"`
	Cookie string `json:"cookie"`
}

type ClientInfoResponseDTO struct {
	ID        uint                    `json:"id"`
	Email     string                  `json:"email"`
	Organize  OrganizerMOD.Controller `json:"organize"`
	Assistant AssistantMOD.Controller `json:"assist"`
	Favs      []EventMOD.Controller   `json:"favs"`
}
