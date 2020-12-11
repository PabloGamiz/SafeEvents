package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txAddFav represents an
type txAddFav struct {
	request  clientDTO.ClientFavDTO
	sessCtrl session.Controller
	ctx      context.Context
}

/*
func (tx *txSignin) buildSessionResponseDTO(ctrl sessionMOD.Controller) *clientDTO.SigninResponseDTO {
	cookie := ctrl.Cookie()
	deadline, _ := ctrl.Deadline() // by sure the session context has a deadline

	return &clientDTO.SigninResponseDTO{
		Cookie:   cookie,
		Deadline: deadline.Unix(),
	}
}
*/
// Precondition validates the transaction is ready to run
func (tx *txAddFav) Precondition() (err error) {
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txAddFav) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a AddFav request for event %d and cookie %s", tx.request.EventID, tx.request.Cookie)

	evnt, err := eventMOD.FindEventByID(ctx, uint(tx.request.EventID))
	if err != nil {
		log.Printf("Error finding Event ID %d", tx.request.EventID)
		return
	}

	var ctr client.Controller
	var ctrID = tx.sessCtrl.GetID()
	if ctr, err = clientMOD.FindClientByID(ctx, ctrID); err != nil {
		return
	}

	ctr.AddFav(evnt.GetEvent()) //CHAPUZA
	clientgw := clientGW.NewClientGateway(ctx, ctr)
	clientgw.AddFavorit()
	return evnt, err
}

// Commit commits the transaction result
func (tx *txAddFav) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txAddFav) Rollback() {

}
