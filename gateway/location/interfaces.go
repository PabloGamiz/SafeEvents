package location

import "github.com/PabloGamiz/SafeEvents-Backend/model/location"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	location.Controller
	Insert() error
}
