package feedback

// RequestDTO represents a feedback request body
type RequestDTO struct {
	ID      uint   `json:"id"`
	Rating  uint   `json:"rating"`
	Message string `json:"message"`
	EventID uint   `json:"eventId"`
	Cookie  string `json:"cookie"`
}
