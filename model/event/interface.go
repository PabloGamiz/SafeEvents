package event

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
)

// Controller represents an Event and its main data.
type Controller interface {
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
	SetLocation(location string)
	GetServices() []service.Controller
	// SetServices(ctrl []service.Controller) int
	GetFeedbacks() []feedback.Controller
}
