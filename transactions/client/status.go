package client

import (
	"context"
	"log"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	interactionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txStatus represents an
type txStatus struct {
	request  clientDTO.StatusRequestDTO
	sessCtrl session.Controller
	ctx      context.Context
}

func (tx *txStatus) buildStatusResponse() *clientDTO.StatusResponseDTO {
	return &clientDTO.StatusResponseDTO{
		DoneAt: time.Now(),
	}
}

func (tx *txStatus) notifyCloseClients() (err error) {
	padding := tx.request.Date.Add(marginTime * -time.Hour)
	var closeTo []uint
	if closeTo, err = interactionMOD.FindCloseToByClientIDAndTime(tx.sessCtrl.GetID(), padding); err != nil {
		return
	}

	for _, clientID := range closeTo {
		log.Printf("HAHA! You are fu**ed boy: %v", clientID)
	}

	return
}

// Precondition validates the transaction is ready to run
func (tx *txStatus) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txStatus) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a status updating request from client %s", tx.sessCtrl.GetEmail())

	if err = tx.sessCtrl.SetStatus(tx.request.Status, tx.request.Date); err != nil {
		return
	}

	if tx.request.Status == clientMOD.POSITIVE {
		if err = tx.notifyCloseClients(); err != nil {
			return
		}
	}

	tx.ctx = ctx
	response := tx.buildStatusResponse()
	return response, nil
}

// Commit commits the transaction result
func (tx *txStatus) Commit() error {
	gw := clientGW.NewClientGateway(tx.ctx, tx.sessCtrl)
	return gw.Update()
}

// Rollback rollbacks any change caused while the transaction
func (tx *txStatus) Rollback() {

}