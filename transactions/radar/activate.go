package radar

import (
	"context"
	"log"
	"time"

	radarDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/radar"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txActivate represents an
type txActivate struct {
	request  radarDTO.ActivateRequestDTO
	sessCtrl session.Controller
}

func (tx *txActivate) buildActivateResponse() *radarDTO.ActivateResponseDTO {
	return &radarDTO.ActivateResponseDTO{
		StartAt: time.Now(),
	}
}

// Precondition validates the transaction is ready to run
func (tx *txActivate) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txActivate) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Radar activation request for client %s", tx.sessCtrl.GetEmail())

	if err = tx.sessCtrl.InitRadar(tx.request.MAC); err != nil {
		return
	}

	response := tx.buildActivateResponse()
	return response, nil
}

// Commit commits the transaction result
func (tx *txActivate) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txActivate) Rollback() {

}
