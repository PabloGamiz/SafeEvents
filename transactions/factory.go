package transactions

import (
	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxSignup builds a brand new transaction for Signup
func NewTxSignup(request clientDTO.SignupRequestDTO) transaction.Tx {
	body := &txSignup{request}
	return transaction.NewTransaction(body)
}
