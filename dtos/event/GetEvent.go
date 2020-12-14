package event

// GetEvent represents the expected data for getting an Event.
type GetEvent struct {
	Cookie string `json:"cookie"`
	ID     uint   `json:"id"`
}
