package event

// Controller represents an Event and its main data.
type Controller interface {
	GetID()
	SetID()
	GetTitle() string
	SetTitle(title string)
	GetDescription() string
	SetDescription(description string)
	GetCapacity() int
	SetCapacity(capacity int)
	GetCheckInDate() int64
	SetCheckInDate(checkInDate int64)
	GetClosureDate() int64
	SetClosureDate(closureDate int64)
	GetLocation()
	SetLocation()
	GetOrganizers()
	GetServices()
	SetServices()
}
