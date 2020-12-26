package radar

import "github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	interaction.Controller
	Insert() error
	Update() error
	Remove() error
}
