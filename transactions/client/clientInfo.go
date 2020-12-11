package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	organizerDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client/organizer"
	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	organizerMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

type txClientInfo struct {
	request  clientDTO.ClientInfoRequestDTO
	sessCtrl session.Controller
}

func (tx *txClientInfo) BuildClientDTO(ctrl clientMOD.Controller) *clientDTO.DTO {
	id := ctrl.GetID()
	email := ctrl.GetEmail()
	organize := tx.BuildOrganizerDTO(ctrl.GetOrganizer())
	return &clientDTO.DTO{
		ID:       id,
		Email:    email,
		Organize: organize,
	}
}

func (tx *txClientInfo) BuildOrganizerDTO(org organizerMOD.Controller) *organizerDTO.DTO {
	id := org.GetID()
	organizes := org.GetEventOrg()
	length := len(organizes)
	ctrls := make([]eventDTO.DTO, length)
	for index, event := range organizes {
		ctrls[index] = tx.BuildEventDTO(event)
	}
	return &organizerDTO.DTO{
		ID:        id,
		Organizes: ctrls,
	}
}

func (tx *txClientInfo) BuildEventDTO(ctrl event.Controller) *eventDTO.DTO {

	return &eventDTO.DTO{
		Title:       ctrl.GetTitle(),
		Description: ctrl.GetDescription(),
		Capacity:    ctrl.GetCapacity(),
		CheckInDate: ctrl.GetCheckInDate(),
		ClosureDate: ctrl.GetClosureDate(),
		Location:    ctrl.GetLocation(),
		Price:       ctrl.GetPrice(),
		Taken:       ctrl.GetTaken(),
		Image:       ctrl.GetImage(),
	}
}

func (tx *txClientInfo) Precondition() error {
	//Comprovar que els elements son correctes
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txClientInfo) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a client info request for client %d", tx.request.ID)

	var sess sessionMOD.Controller
	if sess, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		log.Printf("No session found for provided cookie")
		return
	}

	var id = tx.request.ID

	var ctrl clientMOD.Controller
	if id != 0 {
		if ctrl, err = client.FindClientByID(ctx, id); err != nil {
			return
		}
		response := tx.BuildClientDTO(ctrl)
		return response, err
	} else {
		ctrl = sess.Client()
	}

	return ctrl, err

}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
