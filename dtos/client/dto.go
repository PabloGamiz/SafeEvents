package client

import (
	organizer_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/client/organizer"
)

// DTO represents the expected data from an Event.
type DTO struct {
	ID       uint              `json:"id"`
	Email    string            `json:"email"`
	Organize organizer_api.DTO `json:"organizer"`
}
