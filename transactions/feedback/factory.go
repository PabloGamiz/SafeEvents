package feedback

import (
	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxPOSTFeedback builds a brand new transaction for Posting a feedback
func NewTxPOSTFeedback(request feedbackDTO.RequestDTO) transaction.Tx {
	body := &txPOSTFeedback{request: request}
	return transaction.NewTransaction(body)
}

// NewTxPUTFeedback builds a brand new transaction for editing a feedback
func NewTxPUTFeedback(request feedbackDTO.RequestDTO) transaction.Tx {
	body := &txPUTFeedback{request: request}
	return transaction.NewTransaction(body)
}

// NewTxDELETEFeedback builds a brand new transaction for editing a feedback
func NewTxDELETEFeedback(request feedbackDTO.RequestDTO) transaction.Tx {
	body := &txPUTFeedback{request: request}
	return transaction.NewTransaction(body)
}
