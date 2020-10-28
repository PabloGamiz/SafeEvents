package dtos

// ServiceDTO represents the expected data from a Service.
type ServiceDTO struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Kind        string       `json:"kind"`
	Location    LocationDTO  `json:"location"`
	Product     []ProductDTO `json:"product"`
}
