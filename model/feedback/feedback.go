package feedback

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
)

// Feedback represents the feedback class from UML
type Feedback struct {
	ID          uint                 `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Rating      uint                 `json:"rating" gorm:"not null"`
	Message     string               `json:"message"`
	EventID     uint                 `json:"-" gorm:"index:uc_assistant_event,unique; not null"`
	Assistant   *assistant.Assistant `json:"assistant" gorm:"foreignkey:AssistantID"`
	AssistantID uint                 `json:"-" gorm:"index:uc_assistant_event,unique; not null"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

// GetID return the ID of the Event.
func (feedback *Feedback) GetID() uint {
	return feedback.ID
}

// GetRating return the Name of the Event.
func (feedback *Feedback) GetRating() uint {
	return feedback.Rating
}

// SetRating sets the Name of the Event.
func (feedback *Feedback) SetRating(rating uint) {
	feedback.Rating = rating
}

// GetMessage return the Name of the Event.
func (feedback *Feedback) GetMessage() string {
	return feedback.Message
}

// SetMessage sets the Name of the Event.
func (feedback *Feedback) SetMessage(message string) {
	feedback.Message = message
}

// GetAssistant gets the assistant who has provided the feedback.
func (feedback *Feedback) GetAssistant() assistant.Controller {
	return feedback.Assistant
}

// GetAssistantID gets the assistant who has provided the feedback.
func (feedback *Feedback) GetAssistantID() uint {
	return feedback.AssistantID
}

// GetFeedback gets a pointer to this feedback.
func (feedback *Feedback) GetFeedback() *Feedback {
	return &Feedback{}
}
