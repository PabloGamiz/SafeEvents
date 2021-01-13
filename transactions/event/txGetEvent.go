package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txSignup represents an
type txGetEvent struct {
	request  eventDTO.GetEvent
	faved    bool
	organizs string
	sessCtrl session.Controller
	esOrg    bool
}

func (tx *txGetEvent) buildEventResponseDTO(ctrl event.Controller) *eventDTO.FullEvent {

	return &eventDTO.FullEvent{
		Title:       ctrl.GetTitle(),
		Description: ctrl.GetDescription(),
		Capacity:    ctrl.GetCapacity(),
		CheckInDate: ctrl.GetCheckInDate(),
		ClosureDate: ctrl.GetClosureDate(),
		Location:    ctrl.GetLocation(),
		Price:       ctrl.GetPrice(),
		Taken:       ctrl.GetTaken(),
		Image:       ctrl.GetImage(),
		Tipus:       ctrl.GetTipus(),
		Mesures:     ctrl.GetMesures(),
		Faved:       tx.faved,
		Organizer:   tx.organizs,
		EsOrganize:  tx.esOrg,
	}
}

// Precondition validates the transaction is ready to run
func (tx *txGetEvent) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txGetEvent) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Event request for event with and ID of %d ", tx.request.ID)
	var gw event.Controller
	if gw, err = eventMOD.FindEventByID(tx.request.ID); err != nil {
		return
	}
	var ctr client.Controller
	var ctrID = tx.sessCtrl.GetID()
	if ctr, err = clientMOD.FindClientByID(ctx, ctrID); err != nil {
		return
	}
	clientgw := clientGW.NewClientGateway(ctx, ctr)
	tx.faved, err = clientgw.FindFavorit(gw)
	tx.organizs, err = clientMOD.FindOrganitzersEvent(ctx, tx.request.ID)
	tx.esOrg = false
	if tx.organizs == ctr.GetEmail() {
		tx.esOrg = true
	}
	response := tx.buildEventResponseDTO(gw)
	return response, err
}

// Commit commits the transaction result
func (tx *txGetEvent) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txGetEvent) Rollback() {

}
