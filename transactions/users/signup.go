package users

import (
	"context"
	"log"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
)

// txSignup represents an
type txSignup struct {
	request clientDTO.SignupRequestDTO
}

// Precondition validates the transaction is ready to run
func (tx *txSignup) Precondition() error {
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txSignup) Postcondition(context.Context) (interface{}, error) {
	log.Printf("Got a Signup request")
	response := clientDTO.SignupResponseDTO{
		Cookie:   "hello world :)",
		Deadline: 0,
	}

	return response, nil
}

// Commit commits the transaction result
func (tx *txSignup) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSignup) Rollback() {

}
