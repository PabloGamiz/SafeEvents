package transactions

import (
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxSignup builds a brand new transaction for Signup
func NewTxSignup() transaction.Tx {
	body := &txSignup{}
	return transaction.NewTransaction(body)
}
