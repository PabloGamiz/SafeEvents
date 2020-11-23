package client

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txAddFav represents an
type txAddFav struct {
	request  clientDTO.ClientFavDTO
	sessCtrl session.Controller
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
	log.Printf("Got a AddFav request for event %d and cookie /s", tx.request.EventID, tx.request.Cookie)

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
	return nil, nil
}

// Commit commits the transaction result
func (tx *txAddFav) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txAddFav) Rollback() {

}