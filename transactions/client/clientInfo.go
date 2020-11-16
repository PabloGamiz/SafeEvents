package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type txClientInfo struct {
	request clientDTO.ClientInfoRequestDTO
}

func (tx *txClientInfo) buildClientInfoDTO(client clientMOD.Controller) *clientDTO.ClientInfoResponseDTO {
	//username := client.Username()
	email := client.GetEmail()
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
	log.Printf("Got a Client Info request for client ID %d", tx.request.ID)
	var cli clientMOD.Controller
	if cli, err = clientMOD.GetClientByID(tx.request.ID); err != nil {
		//Do something
		log.Printf("Client exists")
		response := tx.buildClientInfoDTO(cli)
		log.Printf("ResponseDTO created")
		log.Printf("Got a client %d with email %v", tx.request.ID, response.Email)
		return response, nil
	}
	//get events from the client
	response := tx.buildClientInfoDTO(cli)

	return response, nil
}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
