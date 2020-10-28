package location

// Controller represents a Location and it's main data
type Controller interface {
	GetID() string
	SetID(id string)
	GetName() string
	SetName(name string)
	GetAddress() string
	SetAddress(address string)
	GetCoordinates() string
	SetCoordinates(coordinates string)
	GetExtension() int
	SetExtension(extension int)
}
