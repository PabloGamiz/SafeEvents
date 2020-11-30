package assistant

import "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

// Parent is the Client who is being Assistant of some events
type Parent interface {
	GetID() uint
}

// Controller for Assistant
type Controller interface {
	GetID() uint
	AddPurchase(ctrl ticket.Controller) int
	RemovePurchase(ctrl ticket.Controller)
	GetPurchased() []ticket.Controller
	SetParent(Parent)
	GetID() uint

	// S'ha eliminat GetNewPurchased fent així que GetPurchase retorni
	// tant els nous com els antics. En teoria no hauria de suposar cap mena
	// de problema.
}
