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

// txDelFav represents an
type txDelFav struct {
	request  clientDTO.ClientFavDTO
	sessCtrl session.Controller
	ctx      context.Context
}

// Precondition validates the transaction is ready to run
func (tx *txDelFav) Precondition() (err error) {
	//tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txDelFav) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Del request for event %d and cookie %s", tx.request.EventID, tx.request.Cookie)

	// SESSION //
	/*var sess sessionMOD.Controller
	if sess, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		return
	}

	var ctrl client.Controller
	/*if ctrl, err = client.AddFav(ctx, tx.request.EventID, tx.request.Cookie); err != nil {
		log.Printf("Adding to favs EventID %s", tx.request.EventID)
		if err = tx.registerNewClient(ctx); err != nil {
			return
		}
	}

	response := tx.buildSessionResponseDTO(sess)
	//log.Printf("Got a cookie %s for client %v", response.Cookie, sess.GetEmail())
	return sess, ctrl*/
	evnt, err := eventMOD.FindEventByID(ctx, uint(tx.request.EventID))
	if err != nil {
		log.Printf("Error finding Event ID %d", tx.request.EventID)
		return
	}
	var ctr client.Controller
	ctr, err = clientMOD.FindClientByID(tx.ctx, 2)
	ctr.RemoveFav(evnt.GetEvent()) //CHAPUZA
	clientgw := clientGW.NewClientGateway(tx.ctx, ctr)
	err = clientgw.DeleteFavorit(evnt)
	return evnt, err
}

// Commit commits the transaction result
func (tx *txDelFav) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txDelFav) Rollback() {

}