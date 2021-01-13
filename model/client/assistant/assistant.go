package assistant

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// Assistant its a client that assists to events
type Assistant struct {
	ID           uint             `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Purchased    []*ticket.Ticket `json:"purchased" gorm:"foreignkey:AssistantID;"`
	ClientID     uint             `json:"-"`
	newPurchased []ticket.Controller
	parent       Parent
}

// GetID returns the Assistant ID
func (a *Assistant) GetID() uint {
	return a.ID
}

func (a *Assistant) locateTicket(ctrl ticket.Controller) (index int, ok bool) {
	left := len(a.Purchased)
	max := left + len(a.newPurchased)

	for index = 0; index < max; index++ {
		var purchased ticket.Controller
		if index < left {
			purchased = a.Purchased[index]
		} else {
			purchased = a.newPurchased[index-left]
		}

		if ok = purchased.GetID() == ctrl.GetID(); ok {
			return
		}
	}

	return
}

// SetParent sets the Assistant's client
func (a *Assistant) SetParent(p Parent) {
	if a.parent == nil {
		a.parent = p
		a.ClientID = p.GetID()
	}
}

// AddPurchase adds a new ticket purchase for the client
func (a *Assistant) AddPurchase(ctrl ticket.Controller) int {
	a.newPurchased = append(a.newPurchased, ctrl)
	return len(a.Purchased)
}

// RemovePurchase removes the ticket from the client pruchases
func (a *Assistant) RemovePurchase(ctrl ticket.Controller) {
	if len(a.Purchased) == 0 {
		return
	}

	index, exists := a.locateTicket(ctrl)
	if !exists {
		return
	}

	length := len(a.Purchased)
	if index >= length {
		aux := make([]*ticket.Ticket, length-1)

		var padding int
		for n, ticket := range a.Purchased {
			if index == n {
				padding++
				continue
			}

			aux[n-padding] = ticket
		}

		a.Purchased = aux

	} else {
		left := length
		length = len(a.newPurchased)
		aux := make([]ticket.Controller, length-1)

		var padding int
		for n, ticket := range a.newPurchased {
			if index-left == n {
				padding++
				continue
			}

			aux[n-padding] = ticket
		}

		a.newPurchased = aux
	}
}

// GetPurchased returns all currently tickets purchased by the assistant
func (a *Assistant) GetPurchased() (ctrls []ticket.Controller) {
	lengthA := len(a.Purchased)
	lengthB := len(a.newPurchased)
	if lengthA+lengthB == 0 {
		return
	}

	ctrls = make([]ticket.Controller, lengthA+lengthB)
	for index, ticket := range a.Purchased {
		ctrls[index] = ticket
	}

	for index, ticket := range a.newPurchased {
		ctrls[lengthA+index] = ticket
	}

	return
}
