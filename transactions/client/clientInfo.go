package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
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
	var client clientInfoResponseDTO
	if client, err = clientInfoDTO.FindClientByEmail(ctx, tx.info.Email); err != nil {
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
