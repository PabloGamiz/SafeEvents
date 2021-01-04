package radar

import (
	"context"
	"fmt"
	"log"

	radarDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/radar"
	radarGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/radar"
	radarMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar"
	interactionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txInteraction represents an
type txInteraction struct {
	request      radarDTO.InteractionRequestDTO
	sessCtrl     session.Controller
	interactions []interactionMOD.Controller
	ctx          context.Context
}

func (tx *txInteraction) buildInteractionResponse(howmany int) *radarDTO.InteractionResponseDTO {
	return &radarDTO.InteractionResponseDTO{
		HowMany: uint(howmany),
	}
}

// Precondition validates the transaction is ready to run
func (tx *txInteraction) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txInteraction) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got an Interaction request from client %s", tx.sessCtrl.GetEmail())

	if tx.sessCtrl.GetRadar() == nil {
		err = fmt.Errorf("The provided client has no radar running")
		return
	}

	for _, MAC := range tx.request.CloseTo {
		radarCtrl, err := radarMOD.FindRadarByMAC(MAC)
		if err != nil {
			continue
		}

		log.Printf("Got potential infected client, mac = %s", MAC)
		closeToID := radarCtrl.GetID()
		newInteraction := interactionMOD.New(tx.sessCtrl.GetRadar().GetID(), closeToID, tx.request.Instant)
		tx.interactions = append(tx.interactions, newInteraction)
	}

	tx.ctx = ctx
	howMany := tx.sessCtrl.GetRadar().SetInteractions(tx.interactions)
	response := tx.buildInteractionResponse(howMany)
	return response, nil
}

// Commit commits the transaction result
func (tx *txInteraction) Commit() (err error) {
	for _, interact := range tx.interactions {
		gateway := radarGW.NewInteractionGateway(tx.ctx, interact)
		if err = gateway.Insert(); err != nil {
			return
		}
	}

	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txInteraction) Rollback() {
	if size := len(tx.interactions); size > 0 {
		tx.sessCtrl.GetRadar().PopInteractions(size)
	}
}
