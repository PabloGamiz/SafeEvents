package ticket

// PurchaseRequestDTO is the expected struct for a purchase request
type PurchaseRequestDTO struct {
	ClientID    uint   `json:"client_id"`
	EventID     uint   `json:"event_id"`
	Option      uint   `json:"option"`
	HowMany     int    `json:"how_many"`
	Description string `json:"description"`
}

// PurchaseResponseDTO is the response for a purchase request
type PurchaseResponseDTO struct {
	TicketsID []uint `json:"tickets_id"`
}
