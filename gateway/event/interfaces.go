package event

import "github.com/PabloGamiz/SafeEvents-Backend/model/event"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	event.Controller
	Insert() error
}
