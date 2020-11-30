package client

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// Client its the main data object from a client
type Client struct {
	ID       uint                `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Email    string              `json:"email" gorm:"not null; unique"`
	Organize organizer.Organizer `json:"organize" gorm:"foreignKey:ClientID"`
	Assists  assistant.Assistant `json:"assists" gorm:"foreignKey:ClientID"`
	Favs     []*event.Event      `json:"favs" gorm:"many2many:clients_favs;"`
}

// GetID return the id of the client
func (client *Client) GetID() uint {
	return client.ID
}

// GetEmail return the email of the client
func (client *Client) GetEmail() string {
	return client.Email
}

// GetAssistant returns the assistant role of the client
func (client *Client) GetAssistant() assistant.Controller {
	return &client.Assists
}

// GetOrganizer returns the organizer role of the client
func (client *Client) GetOrganizer() organizer.Controller {
	return &client.Organize
}
