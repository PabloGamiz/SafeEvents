package ticket

// PriceTicketRequestDTO is the expected struct for a purchase request
type PriceTicketRequestDTO struct {
	Cookie   string `json:"cookie"`
	TicketID uint   `json:"ticket_id"`
}

// PriceTicketResponseDTO is the response for a purchase request
type PriceTicketResponseDTO struct {
	Price float32 `json:"price"`
}
