package client

import (
	"fmt"
	"time"

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
	Status   Status              `json:"status" gorm:"not null"`
	Updated  time.Time           `json:"update_date"`
}

// SetStatus updates the status of the client
func (client *Client) SetStatus(status Status, update time.Time) error {
	if !client.Updated.Before(update) {
		return fmt.Errorf("The profile is ahead of the provided update time: %v", client.Updated)
	}

	client.Status = status
	client.Updated = update
	return nil
}

// GetStatus return the status of the client
func (client *Client) GetStatus() Status {
	return client.Status
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

// AddFav add a new favourite event to the client
func (client *Client) AddFav(ctrl *event.Event) {
	client.Favs = append(client.Favs, ctrl)
}

// RemoveFav add a new favourite event to the client
func (client *Client) RemoveFav(ctrl *event.Event) {
	index := 0
	for _, i := range client.Favs {
		if i.ID != ctrl.ID {
			client.Favs[index] = i
			index++
		}
	}
}

// GetFavs returns the events that the client has in favourite
func (client *Client) GetFavs() []*event.Event {
	return client.Favs
}
