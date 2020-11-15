package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type txClientInfo struct {
	request clientDTO.clientInfoRequestDTO
}

func (tx *txClientInfo) buildClientInfoDTO(client clientMOD.Controller) *clientDTO.ClientInfoResponseDTO {

	return &clientDTO.ClientInfoDTO{
		Username: username,
		Email:    email,
		Verified: verified,
		Events:   events,
	}
}

func (tx *txClientInfo) Precondition() error {
	//Comprovar que els elements son correctes
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txClientInfo) Postcondition(context.Context) (v interface{}, err error) {
	log.Printf("Got a Client Info request")
	var gw clientGW.Gateway
	if gw, err = clientInfoDTO.FindClientByID(ctx, tx.request.ID); err != nil {
		//Do something
		return
	}
	//get events from the client
	response := tx.buildClientInfoDTO(client)

	return response, nil
}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
