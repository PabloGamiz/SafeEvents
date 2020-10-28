package dtos

// EventDTO represents the expected data from an Event.
type EventDTO struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Capacity    int           `json:"capacity"`
	Organizers  []ClientDTO   `json:"organizers"`
	CheckInDate DateTimeDTO   `json:"checkInDate"`
	ClosureDate DateTimeDTO   `json:"closureDate"`
	Locations   []LocationDTO `json:"locations"`
	Services    []ServiceDTO  `json:"services"`
}
