package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txSignup represents an
type txGetEventAnonim struct {
	request  eventDTO.GetEvent
	faved    bool
	organizs string
	sessCtrl session.Controller
}

func (tx *txGetEventAnonim) buildEventResponseDTO(ctrl event.Controller) *eventDTO.FullEvent {

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
		Organizer:   tx.organizs,
	}
}

// Precondition validates the transaction is ready to run
func (tx *txGetEventAnonim) Precondition() (err error) {
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txGetEventAnonim) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Event request for event with and ID of %d ", tx.request.ID)
	var gw event.Controller
	if gw, err = eventMOD.FindEventByID(tx.request.ID); err != nil {
		return
	}
	tx.organizs, err = clientMOD.FindOrganitzersEvent(ctx, tx.request.ID)
	response := tx.buildEventResponseDTO(gw)
	return response, err
}

// Commit commits the transaction result
func (tx *txGetEventAnonim) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txGetEventAnonim) Rollback() {

}
