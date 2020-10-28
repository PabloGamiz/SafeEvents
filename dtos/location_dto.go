package dtos

// LocationDTO represents the expected data from a Location.
type LocationDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
	Extension   int    `json:"extension"`
}
