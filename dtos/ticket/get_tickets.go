package ticket

// GetTicketsRequestDTO is the expected struct for a purchase request
type GetTicketsRequestDTO struct {
	Cookie  string `json:"cookie"`
	EventID uint   `json:"event_id"`
}
