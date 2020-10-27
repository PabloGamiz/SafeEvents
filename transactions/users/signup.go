package users

import (
	"context"
	"fmt"
	"log"
	"regexp"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txSignup represents an
type txSignup struct {
	request clientDTO.SignupRequestDTO
}

func (tx *txSignup) buildSessionResponseDTO(ctrl session.Controller) *clientDTO.SignupResponseDTO {
	cookie := ctrl.GetCookie()
	deadline := ctrl.GetDeadline()
	return &clientDTO.SignupResponseDTO{
		Cookie:   cookie,
		Deadline: deadline,
	}
}

func (tx *txSignup) registerNewClient() (err error) {
	clnt := &clientMOD.Client{}
	gw := client.NewClientGateway(context.TODO(), clnt)

	return gw.Insert()
}

// Precondition validates the transaction is ready to run
func (tx *txSignup) Precondition() error {
	if matches, err := regexp.MatchString(regexStandardEmail, tx.request.Email); !matches || err != nil {
		return fmt.Errorf(errEmailFormat)
	} else if matches, err := regexp.MatchString(regexBase64, tx.request.Token); !matches || err != nil {
		return fmt.Errorf(errTokenFormat)
	}

	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txSignup) Postcondition(ctx context.Context) (v interface{}, err error) {
	// SESSION //
	if len(tx.request.Token) > 0 {
		log.Printf("Got a Signup request for session %s", tx.request.Token)

		var sess session.Controller
		if sess, err = session.GetSession(tx.request.Token); err == nil {
			log.Printf("The session %s does exists", tx.request.Token)
			response := tx.buildSessionResponseDTO(sess)
			return response, nil
		}
	}

	// SIGNUP //
	var gw client.Gateway
	if gw, err = client.FindClientByEmail(ctx, tx.request.Email); err != nil {
		log.Printf("Signing up a new user %s", tx.request.Email)
		if err = tx.registerNewClient(); err != nil {
			return
		}
	}

	// LOGIN //
	log.Printf("Loging in the user %s", tx.request.Email)
	if gw, err = client.FindClientByEmail(ctx, tx.request.Email); err != nil {
		// At this point the client must be stored in the database
		return
	}

	var sess session.Controller
	if sess, err = session.NewSession(gw); err != nil {
		return
	}

	response := tx.buildSessionResponseDTO(sess)
	return response, nil
}

// Commit commits the transaction result
func (tx *txSignup) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSignup) Rollback() {

}
