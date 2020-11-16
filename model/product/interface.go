package product

// Controller represents a Product and it's main data
type Controller interface {
	GetID() uint
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
	GetPrice() int
	SetPrice(price int)
	GetStatus() Status
	SetStatus(status Status)
}
