package event

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
)

// Controller represents an Event and its main data.
type Controller interface {
	GetEvent() *Event
	GetID() uint
	GetTitle() string
	SetTitle(title string)
	GetDescription() string
	SetDescription(description string)
	GetCapacity() int
	SetCapacity(capacity int)
	GetPrice() float32
	SetPrice(float32)
	GetCheckInDate() time.Time
	SetCheckInDate(checkInDate time.Time)
	GetClosureDate() time.Time
	SetClosureDate(closureDate time.Time)
	GetLocation() string
	SetLocation(string)
	GetServices() []service.Service
	SetServices(services []service.Service)
	TakeTickets(int) error
}
