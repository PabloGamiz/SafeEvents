package client

import (
	"context"
	"log"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txLogout represents an
type txLogout struct {
	request clientDTO.LogoutRequestDTO
}

func (tx *txLogout) buildSessionResponseDTO(ctrl sessionMOD.Controller) *clientDTO.LogoutResponseDTO {
	cookie := ctrl.Cookie()

	return &clientDTO.LogoutResponseDTO{
		Cookie:   cookie,
		Deadline: time.Now().Unix(),
	}
}

// Precondition validates the transaction is ready to run
func (tx *txLogout) Precondition() (err error) {
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txLogout) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Logout request for cookie %s", tx.request.Cookie)

	var sess sessionMOD.Controller
	if sess, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		return
	}

	log.Printf("Loging out session for client %s", sess.GetEmail())
	if err = sessionMOD.KillSession(sess.Cookie()); err != nil {
		return
	}

	response := tx.buildSessionResponseDTO(sess)
	log.Printf("Client %s loged out succesfully", sess.GetEmail())
	return response, nil
}

// Commit commits the transaction result
func (tx *txLogout) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txLogout) Rollback() {

}
