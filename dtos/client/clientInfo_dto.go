package client

import (
	//organizer_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/client/organizer"
	OrganizerMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
)

// ClientInfoRequestDTO represents the expected data from a Client Info request
type ClientInfoRequestDTO struct {
	ID     uint   `json:"id"`
	Cookie string `json:"cookie"`
}

type ClientInfoResponseDTO struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	//Organize organizer_api.DTO `json:"organizer"`
	Organize OrganizerMOD.Controller `json:"organize"`
}
