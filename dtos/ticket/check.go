package ticket

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// CheckRequestDTO is the expected struct for a check request
type CheckRequestDTO struct {
	Cookie  string `json:"cookie"`
	EventID uint   `json:"event_id"`
	Qr      string `json:"qr_code"`
}

// CheckResponseDTO is the response for a check request
type CheckResponseDTO struct {
	Client client.Controller `json:"client"`
	Ticket ticket.Controller `json:"ticket"`
}
