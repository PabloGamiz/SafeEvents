package mail

import "github.com/alvidir/util/pattern/transaction"

// NewTxSendMail builds a brand new transaction for sending a predefined mail to all emails present on request
func NewTxSendMail(request []uint) transaction.Tx {
	body := &txSendMail{request: request}
	return transaction.NewTransaction(body)
}
