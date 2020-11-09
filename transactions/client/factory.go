package client

import (
	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxSignin builds a brand new transaction for Signin
func NewTxSignin(request clientDTO.SigninRequestDTO) transaction.Tx {
	body := &txSignin{request: request}
	return transaction.NewTransaction(body)
}

// NewTxLogout builds a brand new transaction for Signin
func NewTxLogout(request clientDTO.LogoutRequestDTO) transaction.Tx {
	body := &txLogout{request: request}
	return transaction.NewTransaction(body)
}
