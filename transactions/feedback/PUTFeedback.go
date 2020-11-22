package feedback

import (
	"context"
	"fmt"
	"log"

	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	feedbackGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/feedback"
	assistantMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	feedbackMOD "github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPurchase represents an
type txPUTFeedback struct {
	request       feedbackDTO.RequestDTO
	assistantCtrl assistantMOD.Controller
	feedbackGW    feedbackGW.Gateway
	ctx           context.Context
}

func (tx *txPUTFeedback) buildNewFeedbackGW(ctx context.Context) (gw feedbackGW.Gateway, err error) {
	feedback := &feedbackMOD.Feedback{
		ID:          tx.request.ID,
		Rating:      tx.request.Rating,
		Message:     tx.request.Message,
		EventID:     tx.request.EventID,
		AssistantID: tx.assistantCtrl.GetID(),
	}

	gw = feedbackGW.NewFeedbackGateway(ctx, feedback)
	return
}

// Precondition validates the transaction is ready to run
func (tx *txPUTFeedback) Precondition() (err error) {
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

	// Make sure user has provided feedback for this event
	var feedback feedbackMOD.Controller = nil
	feedback, err = feedbackMOD.FindFeedbackByAssistantIDAndEventID(
		int(tx.assistantCtrl.GetID()),
		int(tx.request.EventID))
	if err != nil {
		return
	}
	if feedback == nil {
		err = fmt.Errorf(errUserHasNoFeedbackOnEvent)
		return
	}

	// Make sure the feedback exists and belongs to the Assistant of the session.
	_, err = feedbackMOD.FindFeedbackByIDAndAssistantIDAndEventID(
		int(tx.assistantCtrl.GetID()),
		int(tx.request.EventID),
		int(tx.request.ID))
	if err != nil {
		return
	}

	return

}

// Postcondition declares a new feedback for a certain event
func (tx *txPUTFeedback) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Feedback request from client")

	// Build Feedback Gateway for editing the feedback on the DB
	tx.ctx = ctx
	tx.feedbackGW, err = tx.buildNewFeedbackGW(ctx)

	return
}

// Commit commits the transaction result
func (tx *txPUTFeedback) Commit() (err error) {
	// Insert the new feedback to the DB
	return tx.feedbackGW.Update()
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPUTFeedback) Rollback() {}
