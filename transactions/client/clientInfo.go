package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

type txClientInfo struct {
	request  clientDTO.ClientInfoRequestDTO
	sessCtrl session.Controller
}

func (tx *txClientInfo) BuildClientInfoResponseDTO(ctrl clientMOD.Controller) *clientDTO.ClientInfoResponseDTO {
	id := ctrl.GetID()
	email := ctrl.GetEmail()
	organize := ctrl.GetOrganizer()
	return &clientDTO.ClientInfoResponseDTO{
		ID:       id,
		Email:    email,
		Organize: organize,
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
		response := tx.BuildClientInfoResponseDTO(ctrl)
		return response, err
	} else {
		ctrl = sess.Client()
		var idClient = ctrl.GetID()
		var clientCtrl clientMOD.Controller
		if clientCtrl, err = client.FindClientByID(ctx, idClient); err != nil {
			return
		}
		return clientCtrl, err
	}

}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
