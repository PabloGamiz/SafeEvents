package ticket

// Controller for ticket
type Controller interface {
	Activate() error
	GetID() uint
}
