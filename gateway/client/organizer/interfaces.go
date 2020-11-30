package organizer

import "github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	organizer.Controller
	Insert() error
	Update() error
	Remove() error
}
