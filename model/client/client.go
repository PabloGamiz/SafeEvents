package client

import "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

// Client its the main data object fro a client
type Client struct {
	ID    uint             `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Email string           `json:"email" gorm:"not null"`
	Favs  []*ticket.Ticket `json:"favs" gorm:"many2many:clients_favs;"`
}

// GetID return the id of the client
func (client *Client) GetID() uint {
	return client.ID
}

// GetEmail return the email of the client
func (client *Client) GetEmail() string {
	return client.Email
}

// AddPurchase adds a new ticket purchase for the client
func (client *Client) AddPurchase(ticket Ticket) int {
	client.Purchased = append(client.Purchased, ticket)
	return len(client.Purchased)
}

// RemovePurchase removes the ticket from the client pruchases
func (client *Client) RemovePurchase(ticket Ticket) {
	if len(client.Purchased) == 0 {
		return
	}

	// aux := make([]Ticket, len(client.Purchased)-1)
	// var padding int
	//for index, ticket := range client.Purchased {
	//	if ticket ==
	//}
	//
	//client.Purchased = append(client.Purchased, ticket)
	//return len(client.Purchased)
}
