package client

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

// StatusRequestDTO represents the expected data from a status setting request
type StatusRequestDTO struct {
	Cookie string        `json:"cookie"`
	Status client.Status `json:"status"`
	Date   time.Time     `json:"date"`
}

// StatusResponseDTO represents the resulting response for a status request
type StatusResponseDTO struct {
	DoneAt time.Time `json:"done_at"`
}
