package event

import (
	client_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	location_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/location"
	service_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/service"
)

// DTO represents the expected data from an Event.
type DTO struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Capacity    int                `json:"capacity"`
	Organizers  []client_api.DTO   `json:"organizers"`
	CheckInDate uint64             `json:"checkInDate"`
	ClosureDate uint64             `json:"closureDate"`
	Locations   []location_api.DTO `json:"locations"`
	Services    []service_api.DTO  `json:"services"`
}
