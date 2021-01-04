package interaction

import "time"

// An Interaction is a close contact between assistants
type Interaction struct {
	ID       uint      `json:"id" gorm:"primaryKey; autoIncrement:true"`
	ClientID uint      `json:"client_id" gorm:"not null"`
	CloseTo  uint      `json:"close_to" gorm:"not null"`
	DoneAt   time.Time `json:"done_at" gorm:"not null"`
}

// GetID returns the ID of this interaction
func (interaction *Interaction) GetID() uint {
	return interaction.ID
}

// GetClientID returns the owner of the radar where the interaction is registered
func (interaction *Interaction) GetClientID() uint {
	return interaction.ClientID
}

// GetCloseTo returns the clientID the owner has been close to
func (interaction *Interaction) GetCloseTo() uint {
	return interaction.CloseTo
}

// GetDoneAt return the instant the interaction has been done
func (interaction *Interaction) GetDoneAt() time.Time {
	return interaction.DoneAt
}
