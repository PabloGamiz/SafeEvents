package event

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
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
	GetCheckInDate() time.Time
	SetCheckInDate(checkInDate time.Time)
	GetClosureDate() time.Time
	SetClosureDate(closureDate time.Time)
	GetLocation() location.Location
	SetLocation(location location.Location)
	GetOrganizers() []client.Client
	SetOrganizers(organizers []client.Client)
	GetServices() []service.Service
	SetServices(services []service.Service)
}
