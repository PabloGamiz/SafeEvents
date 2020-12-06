package feedback

import (
	"context"
	"log"

	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	assistantMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	feedbackMOD "github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	sessionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
)

// txPurchase represents an
type txGETFeedbacks struct {
	request feedbackDTO.RequestDTO
	ctx     context.Context
}

func (tx *txGETFeedbacks) BuildFeedbackResponseDTOWithAssistant(feeedbackCtrls []feedbackMOD.Controller, assistantCtrl assistantMOD.Controller) (resp []feedbackDTO.ResponseDTO) {
	resp = make([]feedbackDTO.ResponseDTO, len(feeedbackCtrls))
	for index, feedbackCtrl := range feeedbackCtrls {
		resp[index] = feedbackDTO.ResponseDTO{
			ID:      feedbackCtrl.GetID(),
			Rating:  feedbackCtrl.GetRating(),
			Message: feedbackCtrl.GetMessage(),
			IsOwner: feedbackCtrl.GetAssistantID() == assistantCtrl.GetID(),
		}
	}

	return
}

func (tx *txGETFeedbacks) BuildFeedbackResponseDTOWithoutAssistant(feeedbackCtrls []feedbackMOD.Controller) (resp []feedbackDTO.ResponseDTO) {
	resp = make([]feedbackDTO.ResponseDTO, len(feeedbackCtrls))
	for index, feedbackCtrl := range feeedbackCtrls {
		resp[index] = feedbackDTO.ResponseDTO{
			ID:      feedbackCtrl.GetID(),
			Rating:  feedbackCtrl.GetRating(),
			Message: feedbackCtrl.GetMessage(),
			IsOwner: false,
		}
	}

	return
}

// Precondition validates the transaction is ready to run
func (tx *txGETFeedbacks) Precondition() (err error) {

	// Make sure event exists
	if _, err = eventMOD.FindEventByID(tx.ctx, tx.request.EventID); err != nil {
		return
	}

	return
}

// Postcondition declares a new feedback for a certain event
func (tx *txGETFeedbacks) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a GET Feedbacks request from client")

	// Does the session exist?
	var sessionCtrl sessionMOD.Controller
	var assistantCtrl assistantMOD.Controller
	if sessionCtrl, err = sessionMOD.GetSessionByID(tx.request.Cookie); err != nil {
		assistantCtrl = nil
	} else {
		assistantCtrl = sessionCtrl.GetAssistant()
	}

	tx.ctx = ctx

	var feedbackCtrls []feedback.Controller
	feedbackCtrls, err = feedbackMOD.FindFeedbackByEventID(int(tx.request.EventID))
	if err != nil {
		return
	}

	if assistantCtrl != nil {
		v = tx.BuildFeedbackResponseDTOWithAssistant(feedbackCtrls, assistantCtrl)
	} else {
		v = tx.BuildFeedbackResponseDTOWithoutAssistant(feedbackCtrls)
	}

	return
}

// Commit commits the transaction result
func (tx *txGETFeedbacks) Commit() (err error) {
	// Insert the new feedback to the DB
	return
}

// Rollback rollbacks any change caused while the transaction
func (tx *txGETFeedbacks) Rollback() {

}
