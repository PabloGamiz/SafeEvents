package location

// DTO represents the expected data from a Location.
type DTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
	Extension   int    `json:"extension"`
}
