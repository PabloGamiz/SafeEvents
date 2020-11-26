package client

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	client.Controller
	Insert() error
	Update() error
	Remove() error
	AddFavorit() error
	DeleteFavorit(ctrl event.Controller) error
	FindFavorit(ctrl event.Controller) error
}
