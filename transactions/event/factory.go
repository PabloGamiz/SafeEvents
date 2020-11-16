package event

import (
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxListEvents builds a brand new transaction for List Events
func NewTxListEvents() transaction.Tx {
	body := &txListEvents{}
	return transaction.NewTransaction(body)
}
