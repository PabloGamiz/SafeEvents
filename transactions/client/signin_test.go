package client

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/joho/godotenv"
)

func TestSignin_postcondition(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Got error %s; while loading dotenv", err.Error())
	}

	subject := &txSignin{
		info: newDummyTokenInfo(),
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	response, err := subject.Postcondition(ctx)
	if err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}

	log.Printf("Session response set as %+v", response)
	session.AllInstancesByEmail.Range(func(key interface{}, value interface{}) bool {
		log.Printf("Got session for email %s", key)
		return true
	})

	if _, err := session.GetSessionByEmail(subject.info.Email); err != nil {
		t.Fatalf("Got error %s, while getting session by email %s", err.Error(), subject.info.Email)
	}
}
