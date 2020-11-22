package assistant

import "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

// Parent is the Client who is being Assistant of some events
type Parent interface {
}

// Controller for Assistant
type Controller interface {
	AddPurchase(ctrl ticket.Controller) int
	RemovePurchase(ctrl ticket.Controller)
	GetPurchased() []ticket.Controller
	GetNewPurchased() []ticket.Controller
	SetParent(Parent)
}
