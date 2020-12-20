package event

import (
	"time"

	service_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/service"
)

// DTO represents the expected data from an Event.
type DTO struct {
	Cookie      string `json:"cookie"`
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	//	Organizers  []client_api.DTO  `json:"organizers"`
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
