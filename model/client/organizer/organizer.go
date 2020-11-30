package organizer

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// Organizer its a client that organizes events
type Organizer struct {
	ID        uint           `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Organizes []*event.Event `json:"organizes" gorm:"many2many:organizers_events;"`
	ClientID  uint           `json:"-"`
	parent    Parent
}

// SetParent sets the Organizer's client
func (o *Organizer) SetParent(p Parent) {
	if o.parent == nil {
		o.parent = p
	}
}

// AddEvent adds a new event that organizes the client
func (o *Organizer) AddEvent(ctrl *event.Event) {
	o.Organizes = append(o.Organizes, ctrl)
}

// GetEventOrg returns the events organized
func (o *Organizer) GetEventOrg() []*event.Event {
	return o.Organizes
}

func (o *Organizer) GetID() uint {
	return o.ID
}
