package client

// FavDTO represents the expected data from a AddFav or DelFav request
type FavDTO struct {
	Cookie  string `json:"cookie"`
	EventID int64  `json:"eventid"`
}
