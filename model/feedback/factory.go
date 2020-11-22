package feedback

import (
	"fmt"
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

var once sync.Once

// OpenFeedbackStream opens an stream ensuring the Feedback table does exist
func OpenFeedbackStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while opening stream", err.Error())
		return
	}

	once.Do(func() {
		db.AutoMigrate(&Feedback{})
	})

	return
}

// FindFeedbackByIDAndAssistantIDAndEventID returns, if exists, the feedback provided by assistantID to eventID
func FindFeedbackByIDAndAssistantIDAndEventID(feedbackID int, assistantID int, eventID int) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenFeedbackStream(); err != nil {
		return
	}

	var feedbacks []*Feedback
	db.Find(&feedbacks, queryFindByIDAssistantIDEventID, feedbackID, assistantID, eventID)
	if len(feedbacks) == 0 {
		err = fmt.Errorf(errNoMatchingFeedbackForAssistantAndEvent)
		return
	}

	ctrl = feedbacks[0]
	return
}

// FindFeedbackByAssistantIDAndEventID returns, if exists, the feedback provided by assistantID to eventID
func FindFeedbackByAssistantIDAndEventID(assistantID int, eventID int) (ctrl Controller, err error) {
	var db *gorm.DB
	if db, err = OpenFeedbackStream(); err != nil {
		return
	}

	var feedbacks []*Feedback
	db.Find(&feedbacks, queryFindByAssistantIDEventID, assistantID, eventID)

	ctrl = feedbacks[0]
	return
}

// FindFeedbackByEventID returns, if exists, the feedback corresponding to eventID
func FindFeedbackByEventID(eventID int) (ctrl []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenFeedbackStream(); err != nil {
		return
	}

	var feedbacksMOD []*Feedback
	db.Find(&feedbacksMOD, queryFindByEventID, eventID)

	ctrl = make([]Controller, len(feedbacksMOD))
	for index, feedback := range feedbacksMOD {
		ctrl[index] = feedback
	}

	return
}

// FindFeedbackByAssistantID returns, if exists, the feedback providad by assistantID
func FindFeedbackByAssistantID(assistantID int) (ctrl []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenFeedbackStream(); err != nil {
		return
	}

	var feedbacksMOD []*Feedback
	db.Find(&feedbacksMOD, queryFindByAssistantID, assistantID)

	ctrl = make([]Controller, len(feedbacksMOD))
	for index, feedback := range feedbacksMOD {
		ctrl[index] = feedback
	}

	return
}
