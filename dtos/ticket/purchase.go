package ticket

// PurchaseRequestDTO is the expected struct for a purchase request
type PurchaseRequestDTO struct {
	ClientID uint `json:"client_id"`
	EventID  uint `json:"event_id"`
	Option   uint `json:"option"`
}

// PurchaseResponseDTO is the response for a purchase request
type PurchaseResponseDTO struct {
	TicketID uint `json:"ticket_id"`
}
