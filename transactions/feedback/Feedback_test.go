package feedback

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
// 	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
// 	"github.com/joho/godotenv"
// )

// var eventDTO eventDTO.DTO = &eventDTO.DTO{
// 	ID: 1,
// 	Title:       "Event TEST",
// 	Description: "Event TEST description",
// 	Capacity:    100,
// 	CheckInDate: "2020-12-01T17:01:10Z",
// 	Price:       50,
// 	ClosureDate: "2020-12-01T17:01:10Z",
// 	Location:    "00000000000",
// 	Tipus:       "TEST",
// }

// var feedbackRequestDTO feedbackDTO.RequestDTO = &feedbackDTO.RequestDTO{
// 	ID: 1,
// 	Rating: 5,
// 	Message: "TEST",
// 	EventID: 1,

// }

// func POSTFeedback_postcondition(t *testing.T) {

// }

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
