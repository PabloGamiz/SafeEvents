package organizer

import (
	event_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
)

type DTO struct {
	ID        uint            `json:"id"`
	Organizes []event_api.DTO `json:"organizes"`
}
