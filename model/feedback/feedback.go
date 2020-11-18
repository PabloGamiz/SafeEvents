package feedback

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

// Feedback represents the feedback class from UML
type Feedback struct {
	ID        uint          `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Rating    uint          `json:"rating" gorm:"not null"`
	Message   string        `json:"message"`
	EventID   uint          `json:"-"`
	ClientID  uint          `json:"-"`
	Client    client.Client `json:"client"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
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
