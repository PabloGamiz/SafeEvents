package session

import "github.com/PabloGamiz/SafeEvents-Backend/model/client"

// Controller represents a client and it's main data
type Controller interface {
	client.Controller
	GetCookie() string
	GetDeadline() uint64
}
