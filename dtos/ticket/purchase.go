package ticket

import "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

// PurchaseRequestDTO is the expected struct for a purchase request
type PurchaseRequestDTO struct {
	Cookie      string `json:"cookie"`
	EventID     uint   `json:"event_id"`
	Option      uint   `json:"option"`
	HowMany     int    `json:"how_many"`
	Description string `json:"description"`
}

// PurchaseResponseDTO is the response for a purchase request
type PurchaseResponseDTO struct {
	Tickets []ticket.Controller `json:"tickets_id"`
}
