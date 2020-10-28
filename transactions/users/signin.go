package users

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	"github.com/PabloGamiz/SafeEvents-Backend/google"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"google.golang.org/api/oauth2/v2"
)

// txSignin represents an
type txSignin struct {
	request clientDTO.SigninRequestDTO
	info    *oauth2.Tokeninfo
}

func (tx *txSignin) buildSessionResponseDTO(ctrl session.Controller) *clientDTO.SigninResponseDTO {
	cookie := ctrl.Cookie()
	deadline, _ := ctrl.Deadline() // by sure the session context has a deadline

	return &clientDTO.SigninResponseDTO{
		Cookie:   cookie,
		Deadline: deadline.Unix(),
	}
}

func (tx *txSignin) registerNewClient(ctx context.Context) (err error) {
	clnt := &clientMOD.Client{}
	gw := client.NewClientGateway(ctx, clnt)
	return gw.Insert()
}

// Precondition validates the transaction is ready to run
func (tx *txSignin) Precondition() (err error) {
	tx.info, err = google.VerifyTokenID(tx.request.TokenID)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txSignin) Postcondition(ctx context.Context) (v interface{}, err error) {
	// SESSION //
	if len(tx.request.TokenID) > 0 {
		log.Printf("Got a Signin request for client %s", tx.info.Email)

		var sess session.Controller
		if sess, err = session.GetSessionByEmail(tx.info.Email); err == nil {
			log.Printf("The session for %s does exists", tx.info.Email)
			response := tx.buildSessionResponseDTO(sess)
			return response, nil
		}
	}

	// SIGNUP //
	var gw client.Gateway
	if gw, err = client.FindClientByEmail(ctx, tx.info.Email); err != nil {
		log.Printf("Signing up a new user %s", tx.info.Email)
		if err = tx.registerNewClient(ctx); err != nil {
			return
		}
	}

	// LOGIN //
	log.Printf("Loging in the user %s", tx.info.Email)
	if gw, err = client.FindClientByEmail(ctx, tx.info.Email); err != nil {
		// At this point the client must be stored in the database
		return
	}

	log.Printf("Got a duration for the current session of %v unix", tx.info.ExpiresIn)
	sessCtx := context.TODO()

	var sess session.Controller
	if sess, err = session.NewSession(sessCtx, gw); err != nil {
		return
	}

	response := tx.buildSessionResponseDTO(sess)
	return response, nil
}

// Commit commits the transaction result
func (tx *txSignin) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSignin) Rollback() {

}
