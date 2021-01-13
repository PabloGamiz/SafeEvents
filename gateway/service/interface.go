package service

import "github.com/PabloGamiz/SafeEvents-Backend/model/service"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	service.Controller
	Insert() error
	Update() error
	Remove() error
}
