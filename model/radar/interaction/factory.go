package interaction

import "time"

// New builds a brand new interaction between two clients
func New(owner, closeTo uint, instant time.Time) Controller {
	return &Interaction{
		ClientID: owner,
		CloseTo:  closeTo,
		DoneAt:   instant,
	}
}
