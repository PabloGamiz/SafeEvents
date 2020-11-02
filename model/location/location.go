package location

import "gorm.io/gorm"

// Location represents the Location class from UML.
type Location struct {
	gorm.Model
	ID          uint   `json:"id" sql:"AUTO_INCREMENT"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates" gorm:"primary_key"`
	Extension   int    `json:"extension"`
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
