package event

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// Event represents the Event class from UML.
type Event struct {
	ID          uint              `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Title       string            `json:"title" gorm:"not null;unique"`
	Description string            `json:"description"`
	Capacity    int               `json:"capacity" gorm:"not null"`
	Taken       int               `json:"taken" gorm:"not null"` // How many tickets have been purchased; Capacity - Taken = available_tickets
	Price       float32           `json:"price" gorm:"not null"`
	CheckInDate time.Time         `json:"checkInDate" gorm:"not null"`
	ClosureDate time.Time         `json:"closureDate" gorm:"not null"`
	Location    location.Location `json:"location" gorm:"foreignkey:LocationID;not null"`
	LocationID  uint64            `json:"-"`
	Tickets     []*ticket.Ticket  `json:"tickets" gorm:"foreignkey:EventID"`
	Services    []service.Service `json:"services" gorm:"foreignkey:EventID"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

// GetID return the ID of the Event.
func (event *Event) GetID() uint {
	return event.ID
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

// GetPrice return the price of one ticket for the Event.
func (event *Event) GetPrice() float32 {
	return event.Price
}

// SetPrice sets the price of one ticket for the Event.
func (event *Event) SetPrice(price float32) {
	event.Price = price
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
func (event *Event) GetLocation() location.Location {
	return event.Location
}

// SetLocation sets the Location of the Event.
func (event *Event) SetLocation(location location.Location) {
	event.Location = location
}

// GetServices return the Services of the Event.
func (event *Event) GetServices() []service.Service {
	return event.Services
}

// SetServices sets the Services of the Event.
func (event *Event) SetServices(services []service.Service) {
	event.Services = services
}
