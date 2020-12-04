package feedback

import (
	"context"
	"log"

	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	feedbackGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/feedback"
	assistantMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	feedbackMOD "github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPurchase represents an
type txPOSTFeedback struct {
	request       feedbackDTO.RequestDTO
	assistantCtrl assistantMOD.Controller
	feedbackGW    feedbackGW.Gateway
	ctx           context.Context
}

func (tx *txPOSTFeedback) buildNewFeedbackGW(ctx context.Context) (gw feedbackGW.Gateway, err error) {
	feedback := &feedbackMOD.Feedback{
		Rating:      tx.request.Rating,
		Message:     tx.request.Message,
		EventID:     tx.request.EventID,
		AssistantID: tx.assistantCtrl.GetID(),
	}

	gw = feedbackGW.NewFeedbackGateway(ctx, feedback)
	return
}

// Precondition validates the transaction is ready to run
func (tx *txPOSTFeedback) Precondition() (err error) {
	// Make sure the session exists
	var sessionCtrl sessionMOD.Controller
	if sessionCtrl, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		return
	}
	tx.assistantCtrl = sessionCtrl.GetAssistant()

	// Make sure event exists
	if _, err = eventMOD.FindEventByID(int(tx.request.EventID)); err != nil {
		return
	}

	return
}

// Postcondition declares a new feedback for a certain event
func (tx *txPOSTFeedback) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Feedback request from client")

	// Build Feedback Gateway for inserting the new feedback to the DB
	tx.ctx = ctx
	tx.feedbackGW, err = tx.buildNewFeedbackGW(ctx)

	return
}

// Commit commits the transaction result
func (tx *txPOSTFeedback) Commit() (err error) {
	// Insert the new feedback to the DB
	return tx.feedbackGW.Insert()
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPOSTFeedback) Rollback() {
	tx.feedbackGW.Remove()
}
