package service

import (
	location_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/location"
	product_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/product"
)

// DTO represents the expected data from a Service.
type DTO struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Kind        string            `json:"kind"`
	Location    location_api.DTO  `json:"location"`
	Product     []product_api.DTO `json:"product"`
}
