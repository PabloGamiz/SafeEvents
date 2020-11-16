package location

import (
	"time"
)

// Location represents the Location class from UML.
type Location struct {
	ID          uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Name        string `json:"name" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	Coordinates string `json:"coordinates" gorm:"not null; unique"`
	Extension   int    `json:"extension" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GetID return the ID of the Location.
func (location *Location) GetID() uint {
	return location.ID
}

// GetName return the Name of the Location.
func (location *Location) GetName() string {
	return location.Name
}

// SetName sets the Name of the Location.
func (location *Location) SetName(name string) {
	location.Name = name
}

// GetAddress return the Name of the Location.
func (location *Location) GetAddress() string {
	return location.Address
}

// SetAddress sets the Name of the Location.
func (location *Location) SetAddress(address string) {
	location.Address = address
}

// GetCoordinates return the Name of the Location.
func (location *Location) GetCoordinates() string {
	return location.Coordinates
}

// SetCoordinates sets the Name of the Location.
func (location *Location) SetCoordinates(coordinates string) {
	location.Coordinates = coordinates
}

// GetExtension return the Name of the Location.
func (location *Location) GetExtension() int {
	return location.Extension
}

// SetExtension sets the Name of the Location.
func (location *Location) SetExtension(extension int) {
	location.Extension = extension
}
