package transactions

import (
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxSignup builds a brand new transaction for Signup
func NewTxSignup(username string, email string, token string) transaction.Tx {
	body := &txSignup{username, email, token}
	return transaction.NewTransaction(body)
}
