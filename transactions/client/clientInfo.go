package client

import (
	"context"
	"log"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

//Que es el log?
//Que es la interface?
//Check a la pre?
//Que es el gw
//Que es el return 

type txClientInfo struct {
	request clientDTO.clientInfoDTO
}

func (tx *txSignin) buildClientInfoDTO(client clientMOD.Controller) *clientDTO.ClientInfoDTO {
	username := client.Username() //asi???
	email := client.Email()
	verified := client.Verified()
	events := client.Events()

	return &clientDTO.ClientInfoDTO {
		Username: username,
		Email: email,
		Verified: verified,
		Events: events
	}
}

func (tx *txSignup) Precondition() error {
	//Comprovar que els elements son correctes
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txSignup) Postcondition(context.Context) (v interface{}, err error) {
	log.Printf("Got a Client Info request")
	var client clientInfoDTO
	if client, err = clientInfoDTO.FindClientByEmail(ctx, tx.info.Email); err != nil {
		//Do something
	}
	response := tx.buildClientInfoDTO(client)
	
	return response, nil
}

// Commit commits the transaction result
func (tx *txSignup) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSignup) Rollback() {

}