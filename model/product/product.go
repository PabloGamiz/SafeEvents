package product

// Product represents the product class from UML
type Product struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Price       int    `json:"price" bson:"price,omitempty"`
	Status      Status `json:"status" bson:"status,omitempty"`
}

// GetID return the ID of the Product.
func (product *Product) GetID() string {
	return product.ID
}

// SetID sets a new id to the model
func (product *Product) SetID(id string) { //TODO: El tipus del ID ha de ser primitive.ObjectID o int o string (...) ?
	product.ID = id
}

// GetName return the Name of the Product.
func (product *Product) GetName() string {
	return product.Name
}

// SetName sets the Name of the Product.
func (product *Product) SetName(name string) {
	product.Name = name
}

// GetDescription return the Description of the Product.
func (product *Product) GetDescription() string {
	return product.Description
}

// SetDescription sets the Description of the Product.
func (product *Product) SetDescription(description string) {
	product.Description = description
}

// GetPrice return the Price of the Product.
func (product *Product) GetPrice() int {
	return product.Price
}

// SetPrice sets the Price of the Product.
func (product *Product) SetPrice(price int) {
	product.Price = price
}

// GetStatus return the Status of the Product.
func (product *Product) GetStatus() Status {
	return product.Status
}

// SetStatus sets the Status of the Product.
func (product *Product) SetStatus(status Status) {
	product.Status = status
}
