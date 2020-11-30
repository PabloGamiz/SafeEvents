package assistant

import "github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	assistant.Controller
	Insert() error
	Update() error
	Remove() error
}
