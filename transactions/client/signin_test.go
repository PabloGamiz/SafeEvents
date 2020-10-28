package client

import (
	"context"
	"testing"
	"time"

	"google.golang.org/api/oauth2/v2"
)

func newTestTokenInfo() *oauth2.Tokeninfo {
	return &oauth2.Tokeninfo{
		Email:     "testing@gmail.com",
		ExpiresIn: 1,
		UserId:    "1234",
	}
}

func TestPostcondition(t *testing.T) {
	subject := &txSignin{
		info: newTestTokenInfo(),
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if _, err := subject.Postcondition(ctx); err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}
}
