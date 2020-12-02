package client

// ClientFavDTO represents the expected data from a AddFav or DelFav request
type ClientFavDTO struct {
	Cookie   string `json:"cookie"`
	Deadline int64  `json:"deadline"`
	EventID  int64  `json:"eventid"`
}
