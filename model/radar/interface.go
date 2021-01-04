package radar

import interactionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"

// Controller for ticket
type Controller interface {
	Init() error
	Close() error
	GetID() uint
	GetMAC() string
	SetMAC(string) (string, bool)
	SetInteractions([]interactionMOD.Controller) int
	PopInteractions(int) error
}
