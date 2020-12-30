package client

import (
	"context"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestSendMail_postcondition(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Got error %s; while loading dotenv", err.Error())
	}

	subject := &txSendMail{
		request: []string{"safeevents.sl@gmail.com"},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	_, err := subject.Postcondition(ctx)
	if err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}
}
