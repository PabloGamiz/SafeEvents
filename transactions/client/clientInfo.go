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
	/*
		log.Printf("Got a Client Info request for client ID %d", tx.request.ID)
		var client clientMOD.Controller
		if client, err = clientMOD.GetClientByID(tx.request.ID); err != nil {
			//Client does not exist
			return
		}

		//get events from the client

		log.Printf("Client exists")
		response := tx.buildClientInfoDTO(client)
		log.Printf("Got a client %d with email %v", tx.request.ID, response.Email)

		return response, nil
	*/
	log.Printf("Got a Event request for event %d", tx.request.ID)
	var gw clientGW.Gateway
	gw, err = clientGW.FindClientByID(ctx, tx.request.ID)
	return gw, err
}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
