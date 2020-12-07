package feedback

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
)

// Controller represents an Event and its main data.
type Controller interface {
	GetID() uint
	GetRating() uint
	SetRating(rating uint)
	GetMessage() string
	SetMessage(message string)
	GetAssistant() assistant.Controller
	GetAssistantID() uint
	GetFeedback() *Feedback
}
