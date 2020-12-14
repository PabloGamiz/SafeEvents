package ticket

import (
	"context"
	"testing"
	"time"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	clientTX "github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
	eventTX "github.com/PabloGamiz/SafeEvents-Backend/transactions/event"
	"github.com/joho/godotenv"
)

func getDummySession(t *testing.T) session.Controller {
	sess, err := clientTX.SetupDummyUser()
	if err != nil {
		t.Fatalf("Got error %s, while getting dummy session", err.Error())
	}

	return sess
}

func getDummyEventID(ctx context.Context, t *testing.T, cookie string) uint {
	title, err := session.NewSessionID()
	if err != nil {
		t.Fatalf(err.Error())
	}

	request := eventDTO.PublicaEvent{
		Cookie:      cookie,
		Title:       "dummy-" + title,
		Description: "Its an event for example",
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

func TestPurchase_booked(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Got error %s, while loading dotenv", err.Error())
	}

	sess := getDummySession(t)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	option := ticketMOD.BOOKED
	howmany := 1
	cookie := sess.Cookie()
	eventID := getDummyEventID(ctx, t, cookie)
	request := &txPurchase{
		request: ticketDTO.PurchaseRequestDTO{
			Cookie:      cookie,
			EventID:     eventID,
			Option:      uint(option),
			HowMany:     howmany,
			Description: "testing",
		},
		sessCtrl: sess,
	}

	response, err := request.Postcondition(ctx)
	if err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}

	dto, ok := response.(*ticketDTO.PurchaseResponseDTO)
	if !ok {
		t.Fatalf("On response dto assertion")
	}

	if got := len(dto.Tickets); got != howmany || got == 0 {
		t.Fatalf("Got %v purchased tickets, want %v", got, howmany)
	}

	ticketID := dto.Tickets[0].GetID()
	subject, err := ticketMOD.GetTicketByID(ticketID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if got := subject.GetOption(); got != option {
		t.Errorf("Got %v option, want %v", got, option)
	}

	if got := subject.GetEventID(); got != eventID {
		t.Errorf("Got event_id=%v, want %v", got, eventID)
	}

	if got := subject.GetClientID(); got != sess.GetID() {
		t.Errorf("Got event_id=%v, want %v", got, sess.GetID())
	}

	if got := subject.GetQR(); len(got) > 0 {
		t.Errorf("Got unexpected qr_code=%v", got)
	}
}

func TestPurchase_bought(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Got error %s, while loading dotenv", err.Error())
	}

	sess := getDummySession(t)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	option := ticketMOD.BOUGHT
	howmany := 1
	cookie := sess.Cookie()
	eventID := getDummyEventID(ctx, t, cookie)
	request := &txPurchase{
		request: ticketDTO.PurchaseRequestDTO{
			Cookie:      cookie,
			EventID:     eventID,
			Option:      uint(option),
			HowMany:     howmany,
			Description: "testing",
		},
		sessCtrl: sess,
	}

	response, err := request.Postcondition(ctx)
	if err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}

	dto, ok := response.(*ticketDTO.PurchaseResponseDTO)
	if !ok {
		t.Fatalf("On response dto assertion")
	}

	if got := len(dto.Tickets); got != howmany || got == 0 {
		t.Fatalf("Got %v purchased tickets, want %v", got, howmany)
	}

	ticketID := dto.Tickets[0].GetID()
	subject, err := ticketMOD.GetTicketByID(ticketID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if got := subject.GetOption(); got != option {
		t.Errorf("Got %v option, want %v", got, option)
	}

	if got := subject.GetEventID(); got != eventID {
		t.Errorf("Got event_id=%v, want %v", got, eventID)
	}

	if got := subject.GetClientID(); got != sess.GetID() {
		t.Errorf("Got event_id=%v, want %v", got, sess.GetID())
	}

	if got := subject.GetQR(); len(got) == 0 {
		t.Errorf("Got no qr_code")
	}
}
