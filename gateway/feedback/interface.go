package feedback

import "github.com/PabloGamiz/SafeEvents-Backend/model/feedback"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	feedback.Controller
	Insert() error
	Update() error
	Remove() error
}
