package ticket

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"gorm.io/gorm"
)

// Ticket its the main data object fro a Ticket
type Ticket struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Description string         `json:"description" gorm:"not null"`
	Price       float32        `json:"price" gorm:"not null"`
	Client      *client.Client `json:"-" gorm:"not null"`
	ClientID    uint           `json:"-" gorm:"not null"`
	Event       *event.Event   `json:"-" gorm:"not null"`
	EventID     uint           `json:"-" gorm:"not null"`
}

// GetID return the id of the ticket
func (ticket *Ticket) GetID() uint {
	return ticket.ID
}
