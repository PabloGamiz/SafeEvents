package organizer

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// Organizer its a client that organizes events
type Organizer struct {
	ID       uint           `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Organize []*event.Event `json:"organize" gorm:"many2many:organizers_events;"`
	ClientID uint           `json:"-"`
	parent   Parent
}

// SetParent sets the Organizer's client
func (o *Organizer) SetParent(p Parent) {
	if o.parent == nil {
		o.parent = p
	}
}

// AddEvent adds a new event that organizes the client
func (o *Organizer) AddEvent(ctrl *event.Event) {
	o.Organize = append(o.Organize, ctrl)
}

// GetEventOrg returns the event organized
func (o *Organizer) GetEventOrg() []*event.Event {
	return o.Organize
}
