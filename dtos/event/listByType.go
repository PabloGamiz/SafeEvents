package event

// ListEventsByTypeRequestDTO represents the expected data from a List by type request
type ListEventsByTypeRequestDTO struct {
	EventType string `json:"eventType"`
}
