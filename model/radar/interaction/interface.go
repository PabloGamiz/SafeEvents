package interaction

import "time"

// Controller for interaction
type Controller interface {
	GetID() uint
	GetClientID() uint
	GetCloseTo() uint
	GetDoneAt() time.Time
}
