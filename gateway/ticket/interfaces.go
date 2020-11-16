package ticket

import "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	ticket.Controller
	Insert() error
	Update() error
	Remove() error
}
