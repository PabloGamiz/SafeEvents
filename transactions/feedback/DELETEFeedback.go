package feedback

import (
	"context"
	"log"

	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	feedbackGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/feedback"
	assistantMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	feedbackMOD "github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPurchase represents an
type txDELETEFeedback struct {
	request          feedbackDTO.RequestDTO
	assistantCtrl    assistantMOD.Controller
	feedbackToDelete feedbackMOD.Controller
	feedbackGW       feedbackGW.Gateway
	ctx              context.Context
}

// Precondition validates the transaction is ready to run
func (tx *txDELETEFeedback) Precondition() (err error) {
	// Make sure the session exists
	var sessionCtrl sessionMOD.Controller
	if sessionCtrl, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		return
	}
	tx.assistantCtrl = sessionCtrl.GetAssistant()

	// Make sure the feedback exists and belongs to the Assistant of the session.
	tx.feedbackToDelete, err = feedbackMOD.FindFeedbackByIDAndAssistantIDAndEventID(
		int(tx.request.ID),
		int(tx.assistantCtrl.GetID()),
		int(tx.request.EventID))
	if err != nil {
		return
	}

	return

}

// Postcondition declares a new feedback for a certain event
func (tx *txDELETEFeedback) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a DELETE Feedback request from client")

	// Build Feedback Gateway for editing the feedback on the DB
	tx.ctx = ctx
	tx.feedbackGW = feedbackGW.NewFeedbackGateway(ctx, tx.feedbackToDelete)

	return
}

// Commit commits the transaction result
func (tx *txDELETEFeedback) Commit() (err error) {
	// Insert the new feedback to the DB
	return tx.feedbackGW.Remove()
}

// Rollback rollbacks any change caused while the transaction
func (tx *txDELETEFeedback) Rollback() {}
