package ticket

import "time"

// Controller for ticket
type Controller interface {
	Activate() error
	GetID() uint
	GetOption() Option
	GetCreatedAt() time.Time
	GetClientID() uint
	GetEventID() uint
	GetQR() string
	Check() error
}
