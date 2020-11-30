package ticket

// ActivateRequestDTO is the expected struct for a purchase request
type ActivateRequestDTO struct {
	Cookie  string `json:"cookie"`
	EventID uint   `json:"event_id"`
	HowMany int    `json:"how_many"`
}
