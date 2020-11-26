package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

type txClientInfo struct {
	request clientDTO.ClientInfoRequestDTO
}

/*
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
*/

func (tx *txClientInfo) Precondition() error {
	//Comprovar que els elements son correctes
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txClientInfo) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a client info request for client %d", tx.request.ID)
	var ctrl clientMOD.Controller
	if ctrl, err = client.FindClientByID(ctx, tx.request.ID); err != nil {
		return
	}
	/*
		var events eventMOD.Controller
		if events, err = eventGW.FindEventsByClientID(ctx, tx.request.ID); err != nil {
			return
		}
	*/
	//response := tx.buildClientInfoDTO(ctrl)
	//return response, err
	return ctrl, err

}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
