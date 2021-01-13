package event

import (
	"time"

	service_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/service"
)

// PublicaEvent represents the expected data for creating an Event.
type PublicaEvent struct {
	Cookie      string            `json:"cookie"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Capacity    int               `json:"capacity"`
	Organizer   string            `json:"organizers"`
	CheckInDate time.Time         `json:"checkInDate"`
	ClosureDate time.Time         `json:"closureDate"`
	Price       float32           `json:"price"`
	Location    string            `json:"location"`
	Services    []service_api.DTO `json:"services"`
	Image       string            `json:"image"`
	Tipus       string            `json:"tipus"`
	Faved       bool              `json:"Faved"`
	Taken       int               `json:"Taken"`
	Mesures     string            `json:"mesures"`
}
