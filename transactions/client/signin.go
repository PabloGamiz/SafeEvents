package client

import (
	"context"
	"log"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"google.golang.org/api/oauth2/v2"
)

// txSignin represents an
type txSignin struct {
	request clientDTO.SigninRequestDTO
	info    *oauth2.Tokeninfo
}

func (tx *txSignin) buildSessionResponseDTO(ctrl sessionMOD.Controller) *clientDTO.SigninResponseDTO {
	cookie := ctrl.Cookie()
	deadline, _ := ctrl.Deadline() // by sure the session context has a deadline

	return &clientDTO.SigninResponseDTO{
		Cookie:   cookie,
		Deadline: deadline.Unix(),
	}
}

func (tx *txSignin) registerNewClient(ctx context.Context) (err error) {
	clnt := &clientMOD.Client{
		Email: tx.info.Email,
	}

	gw := clientGW.NewClientGateway(ctx, clnt)
	return gw.Insert()
}

// Precondition validates the transaction is ready to run
func (tx *txSignin) Precondition() (err error) {
	//tx.info, err = google.VerifyTokenID(tx.request.TokenID)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txSignin) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Signin request for client" /*tx.info.Email*/)

	// SESSION //
	var sess sessionMOD.Controller
	if sess, err = sessionMOD.GetSessionByEmail(tx.info.Email); err == nil {
		log.Printf("The session for %s already exists", tx.info.Email)
		response := tx.buildSessionResponseDTO(sess)
		return response, nil
	}

	// SIGNUP //
	var gw clientGW.Gateway
	if gw, err = clientGW.FindClientByEmail(ctx, tx.info.Email); err != nil {
		log.Printf("Signing up a new client %s", tx.info.Email)
		if err = tx.registerNewClient(ctx); err != nil {
			return
		}
	}

	// LOGIN //
	log.Printf("Loging in the client %s", tx.info.Email)
	if gw, err = clientGW.FindClientByEmail(ctx, tx.info.Email); err != nil {
		// At this point the client must be stored in the database
		return
	}

	log.Printf("Building session for client %s", gw.GetEmail())
	deadline := time.Unix(tx.info.ExpiresIn, 0)
	sessCtx, _ := context.WithDeadline(context.TODO(), deadline)
	if sess, err = sessionMOD.NewSession(sessCtx, gw); err != nil {
		return
	}

	response := tx.buildSessionResponseDTO(sess)
	log.Printf("Got a cookie %s for client %v", response.Cookie, sess.GetEmail())
	return response, nil
}

// Commit commits the transaction result
func (tx *txSignin) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSignin) Rollback() {

}
