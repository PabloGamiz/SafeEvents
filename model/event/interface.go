package event

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
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
	GetTaken() int
	GetCheckInDate() time.Time
	SetCheckInDate(checkInDate time.Time)
	GetImage() string
	SetImage(Image string)
	GetTipus() string
	SetTipus(tipus string)
	GetClosureDate() time.Time
	SetClosureDate(closureDate time.Time)
	GetLocation() string
	SetLocation(location string)
	GetServices() []service.Controller
	TakeTickets(int) error
	GetFeedbacks() []feedback.Controller
}
