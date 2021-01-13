package event

import (
	"context"
	"fmt"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	serviceDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/service"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	serviceGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/service"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	productMOD "github.com/PabloGamiz/SafeEvents-Backend/model/product"
	serviceMOD "github.com/PabloGamiz/SafeEvents-Backend/model/service"
	sessMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPublicaEvent represents an
type txModificaEvent struct {
	request    eventDTO.DTO
	sessCtrl   sessMOD.Controller
	clientCtrl clientMOD.Controller
	eventCtrl  eventMOD.Controller
}

func (tx *txModificaEvent) GetServicesFromRequest(servicesRequest []serviceDTO.DTO) []*serviceMOD.Service {
	var services = make([]*serviceMOD.Service, len(servicesRequest))
	for index, service := range servicesRequest {
		var serviceProducts = make([]productMOD.Product, len(service.Product))
		for i, product := range service.Product {
			var productCtrl = productMOD.Product{
				ID:          uint(product.ID),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Status:      product.Status,
			}
			serviceProducts[i] = productCtrl
		}
		var serviceCtrl = &serviceMOD.Service{
			ID:          uint(service.ID),
			Name:        service.Name,
			Description: service.Description,
			Kind:        service.Kind,
			Location:    service.Location,
			Products:    serviceProducts,
		}
		services[index] = serviceCtrl
	}

	return services
}

func (tx *txModificaEvent) Precondition() (err error) { //Comprova que no existeix l'event
	// make sure the session exists
	tx.sessCtrl, err = sessMOD.GetSessionByID(tx.request.Cookie)
	if err != nil {
		return
	}

	// Make sure event exists before modifiying.
	if _, err = eventMOD.FindEventByID(uint(tx.request.ID)); err != nil {
		return
	}

	// Check that the autenthicated user is the organizer of the event.
	var organizerCtrl = tx.sessCtrl.GetOrganizer()
	var isOrganizer = false
	for _, event := range organizerCtrl.GetEventOrg() {
		var eventCtrl eventMOD.Controller = event
		if eventCtrl.GetID() == tx.request.ID {
			tx.eventCtrl = eventCtrl
			isOrganizer = true
			break
		}
	}
	if !isOrganizer {
		return fmt.Errorf(errUserIsNotOrganizer)
	}

	return
}

func (tx *txModificaEvent) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Modifica Event request for event %s", tx.request.Title)

	var services = tx.eventCtrl.GetServices()
	for _, service := range services {
		gw := serviceGW.NewServiceGateway(ctx, service)
		if err = gw.Remove(); err != nil {
			return
		}
	}

	tx.eventCtrl.SetTitle(tx.request.Title)
	tx.eventCtrl.SetDescription(tx.request.Description)
	tx.eventCtrl.SetCapacity(tx.request.Capacity)
	tx.eventCtrl.SetPrice(tx.request.Price)
	tx.eventCtrl.SetCheckInDate(tx.request.CheckInDate)
	tx.eventCtrl.SetClosureDate(tx.request.ClosureDate)
	tx.eventCtrl.SetLocation(tx.request.Location)
	tx.eventCtrl.SetServices(tx.GetServicesFromRequest(tx.request.Services))
	tx.eventCtrl.SetImage(tx.request.Image)
	tx.eventCtrl.SetTipus(tx.request.Tipus)
	tx.eventCtrl.SetMesures(tx.request.Mesures)

	gw := eventGW.NewEventGateway(ctx, tx.eventCtrl)
	if err = gw.Update(); err != nil {
		return
	}

	return nil, err
}

func (tx *txModificaEvent) Commit() (err error) {
	return
}

func (tx *txModificaEvent) Rollback() {

}
