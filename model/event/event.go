package event

import (
	"fmt"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
)

// Event represents the Event class from UML.
type Event struct {
	ID          uint                 `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Title       string               `json:"title" gorm:"not null;unique"`
	Description string               `json:"description"`
	Capacity    int                  `json:"capacity" gorm:"not null"`
	Taken       int                  `json:"taken" gorm:"not null;check:,taken <= capacity"` // How many tickets have been purchased; Capacity - Taken = available_tickets
	Price       float32              `json:"price" gorm:"not null"`
	CheckInDate time.Time            `json:"checkInDate" gorm:"not null"`
	ClosureDate time.Time            `json:"closureDate" gorm:"not null"`
	Location    string               `json:"location" gorm:"not null"`
	Feedbacks   []*feedback.Feedback `json:"feedbacks" gorm:"foreignkey:EventID"`
	Services    []*service.Service   `json:"services" gorm:"foreignkey:EventID"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	Image       string               `json:"image" gorm:"not null"`
	Tipus       string               `json:"tipus" gorm:"not null"`
	//mu          sync.Mutex
}

//CHAPUZA
func (event *Event) GetEvent() *Event {
	return event
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

// GetTaken return the number of tickets taken of the Event.
func (event *Event) GetTaken() int {
	return event.Taken
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
func (event *Event) GetLocation() string {
	return event.Location
}

// SetLocation sets the Location of the Event.
func (event *Event) SetLocation(loc string) {
	event.Location = loc
}

// GetServices return the Services of the Event.
func (event *Event) GetServices() (ctrls []service.Controller) {
	length := len(event.Services)
	if length == 0 {
		return
	}

	ctrls = make([]service.Controller, length)
	for index, service := range event.Services {
		ctrls[index] = service
	}

	return
}

// GetFeedbacks return the Feedbacks of the Event.
func (event *Event) GetFeedbacks() (ctrls []feedback.Controller) {
	length := len(event.Feedbacks)
	if length == 0 {
		return
	}

	ctrls = make([]feedback.Controller, length)
	for index, feedback := range event.Feedbacks {
		ctrls[index] = feedback
	}

	return
}

// GetImage return the path of the Image.
func (event *Event) GetImage() string {
	return event.Image
}

// SetImage sets the path of the Image.
func (event *Event) SetImage(image string) {
	event.Image = image
}

// GetTipus return the type of the event.
func (event *Event) GetTipus() string {
	return event.Tipus
}

// SetTipus sets the type of the event.
func (event *Event) SetTipus(tipus string) {
	event.Tipus = tipus
}

// TakeTickets takes as much tickets as set on n, if there is not enought capacity an error its thrown
func (event *Event) TakeTickets(n int) error {
	//event.mu.Lock()
	//defer event.mu.Unlock()

	if event.Taken+n > event.Capacity {
		return fmt.Errorf("Event capacity exceed")
	}

	event.Taken += n
	return nil
}
