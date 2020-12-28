package radar

import (
	"context"
	"log"
	"time"

	radarDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/radar"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txDeactivate represents an
type txDeactivate struct {
	request  radarDTO.ActivateRequestDTO
	sessCtrl session.Controller
}

func (tx *txDeactivate) buildActivateResponse() *radarDTO.ActivateResponseDTO {
	return &radarDTO.ActivateResponseDTO{
		DoneAt: time.Now(),
	}
}

// Precondition validates the transaction is ready to run
func (tx *txDeactivate) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txDeactivate) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Radar deactivation request for client %s", tx.sessCtrl.GetEmail())

	if err = tx.sessCtrl.FinishRadar(); err != nil {
		return
	}

	response := tx.buildActivateResponse()
	return response, nil
}

// Commit commits the transaction result
func (tx *txDeactivate) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txDeactivate) Rollback() {

}
