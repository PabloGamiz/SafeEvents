package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type txClientInfo struct {
	request clientDTO.ClientInfoRequestDTO
}

func (tx *txClientInfo) buildClientInfoDTO(ctrl clientMOD.Controller) *clientDTO.ClientInfoResponseDTO {
	//username := client.Username()
	email := ctrl.GetEmail()
	//verified := client.Verified()
	//events := client.Events()

	return &clientDTO.ClientInfoResponseDTO{
		//Username: username,
		Email: email,
		//Verified: verified,
		//Events:   events,
	}
}

func (tx *txClientInfo) Precondition() error {
	//Comprovar que els elements son correctes
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txClientInfo) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Event request for client %d", tx.request.ID)
	var ctrl clientMOD.Controller
	if ctrl, err = clientGW.FindClientByID(ctx, tx.request.ID); err != nil {
		return
	}
	response := tx.buildClientInfoDTO(ctrl)
	return response, err

}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
