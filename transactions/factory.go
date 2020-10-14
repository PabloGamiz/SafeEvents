package transactions

import (
	"github.com/alvidir/util/pattern/transaction"
)

func newTxSignup() transaction.Tx {
	body := &txSignup{}
	return transaction.NewTransaction(body)
}
