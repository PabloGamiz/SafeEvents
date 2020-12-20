package feedback

import (
	"context"
	"testing"
	"time"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	feedbackMOD "github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	sessMOD "github.com/PabloGamiz/SafeEvents-Backend/model/session"
	clientTX "github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
	eventTX "github.com/PabloGamiz/SafeEvents-Backend/transactions/event"

	"github.com/joho/godotenv"
)

func getDummySession(t *testing.T) sessMOD.Controller {
	sess, err := clientTX.SetupDummyUser()
	if err != nil {
		t.Fatalf("Got error %s, while getting dummy session", err.Error())
	}

	return sess
}

func getDummyEventID(ctx context.Context, t *testing.T, cookie string) uint {
	title, err := sessMOD.NewSessionID()
	if err != nil {
		t.Fatalf(err.Error())
	}

	request := eventDTO.PublicaEvent{
		Cookie:      cookie,
		Title:       "FEEDBACK TEST" + title,
		Description: "Its a test Event",
		Capacity:    15000,
		CheckInDate: time.Now(),
		ClosureDate: time.Now(),
		Price:       120,
		Location:    "Testingcity",
	}

	txPublicaEvent := eventTX.NewTxPublicaEvent(request)
	txPublicaEvent.Execute(ctx)

	content, err := txPublicaEvent.Result()
	if err != nil {
		t.Fatalf(err.Error())
	}

	event, ok := content.(eventMOD.Controller)
	if !ok {
		t.Fatalf("Assertion fail")
	}

	return event.GetID()
}

// POSTFeedbackPostcondition tests the POSTFeedback endpoint
func Postcondition(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Got error %s, while loading dotenv", err.Error())
	}

	sess := getDummySession(t)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	cookie := sess.Cookie()
	eventID := getDummyEventID(ctx, t, cookie)
	request := feedbackDTO.RequestDTO{
		ID:      1,
		Rating:  5,
		Message: "TEST",
		EventID: eventID,
		Cookie:  cookie,
	}

	txPOSTFeedback := NewTxPOSTFeedback(request)
	_, err := txPOSTFeedback.Result()
	if err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}

	_, err = eventMOD.FindEventByID(eventID)
	if err != nil {
		t.Fatalf("Got error %s; while obtaining event related to feedback", err.Error())
	}

	_, err = feedbackMOD.FindFeedbackByAssistantIDAndEventID(int(sess.GetID()), int(eventID))
	if err != nil {
		t.Fatalf("Got error %s; while obtaining feedback", err.Error())
	}

}

// func DELETEFeedback_postcondition(t *testing.T) {
// 	if err := godotenv.Load("../../.env"); err != nil {
// 		t.Fatalf("Got error %s; while loading dotenv", err.Error())
// 	}

// 	var feedbackRequestDTO feedbackDTO.RequestDTO

// 	subject := feedback.NewTxDELETEFeedback(feedbackRequestDTO)

// 	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
// 	defer cancel()

// 	_, err := subject.Postcondition(ctx)
// 	if err != nil {
// 		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
// 	}
// }
