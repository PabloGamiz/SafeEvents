package event

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
)

// Event represents the Event class from UML.
type Event struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Title       string               `json:"title" bson:"title,omitempty"`
	Description string               `json:"description" bson:"description,omitempty"`
	Capacity    int                  `json:"capacity" bson:"capacity,omitempty"`
	CheckInDate time.Time            `json:"checkInDate" bson:"checkInDate,omitempty"`
	ClosureDate time.Time            `json:"closureDate" bson:"closureDate,omitempty"`
	Location    location.Controller  `json:"location" bson:"location"`
	Organizers  []client.Controller  `json:"organizers" bson:"organizers"`
	Services    []service.Controller `json:"services" bson:"services"`
}

// GetID return the ID of the Event.
func (event *Event) GetID() {
}

// SetID sets the Name of the Event.
func (event *Event) SetID() {

}

// GetTitle return the Name of the Event.
func (event *Event) GetTitle() string {
	return event.Title
}

// SetTitle sets the Name of the Event.
func (event *Event) SetTitle(title string) {
	event.Title = title
}

// GetDescription return the Description of the Event.
func (event *Event) GetDescription() string {
	return event.Description
}

// SetDescription sets the Description of the Event.
func (event *Event) SetDescription(description string) {
	event.Description = description
}

// GetCapacity return the Capacity of the Event.
func (event *Event) GetCapacity() int {
	return event.Capacity
}

// SetCapacity sets the Capacity of the Event.
func (event *Event) SetCapacity(capacity int) {
	event.Capacity = capacity
}

// GetCheckInDate return the ChekInDate of the Event.
func (event *Event) GetCheckInDate() time.Time {
	return event.CheckInDate
}

// SetCheckInDate sets the CheckInDate of the Event.
func (event *Event) SetCheckInDate(checkInDate time.Time) {
	event.CheckInDate = checkInDate
}

// GetClosureDate return the ClosureDate of the Event.
func (event *Event) GetClosureDate() time.Time {
	return event.ClosureDate
}

// SetClosureDate sets the CheckInDate of the Event.
func (event *Event) SetClosureDate(closureDate time.Time) {
	event.ClosureDate = closureDate
}

// GetLocation return the Location of the Event.
func (event *Event) GetLocation() location.Controller {
	return event.Location
}

// SetLocation sets the Location of the Event.
func (event *Event) SetLocation(location location.Controller) {
	event.Location = location
}

// GetOrganizers return the Organizers of the Event.
func (event *Event) GetOrganizers() []client.Controller {
	return event.Organizers
}

// GetServices return the Services of the Event.
func (event *Event) GetServices() time.Time {
	return event.CheckInDate
}

// SetServices sets the Services of the Event.
func (event *Event) SetServices(services []service.Controller) {
	event.Services = services
}
