package users

import (
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxSignup builds a brand new transaction for Signup
func NewTxListEvents() transaction.Tx {
	body := &txListEvents{}
	return transaction.NewTransaction(body)
}
