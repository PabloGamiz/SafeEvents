package location

// Location represents the Location class from UML.
type Location struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name,omitempty"`
	Address     string `json:"address" bson:"address,omitempty"`
	Coordinates string `json:"coordinates" bson:"coordinates,omitempty"`
	Extension   int    `json:"extension" bson:"extension,omitempty"`
}

// GetID return the ID of the Location.
func (location *Location) GetID() string {
	return location.ID
}

// SetID sets the Name of the Location.
func (location *Location) SetID(id string) {
	location.ID = id
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
