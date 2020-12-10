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
	} else {
		ctrl = sess.Client()
	}

	return ctrl, err

}

// Commit commits the transaction result
func (tx *txClientInfo) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txClientInfo) Rollback() {

}
