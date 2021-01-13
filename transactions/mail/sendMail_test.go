package mail

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

	request := []uint{1}
	subject := NewTxSendMail(request)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	subject.Execute(ctx)
	if _, err := subject.Result(); err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}
}
