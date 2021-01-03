package client

import (
	"context"
	"fmt"
	"os"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"google.golang.org/api/oauth2/v2"
)

func newDummyTokenInfo() *oauth2.Tokeninfo {
	email, exists := os.LookupEnv(envDummyEmail)
	if !exists {
		email = "safeevents.sl@gmail.com"
	}

	return &oauth2.Tokeninfo{
		Email:     email,
		ExpiresIn: -1,
		UserId:    "1234",
	}
}

// SetupDummyUser inits a dummy user for testing
func SetupDummyUser() (sess session.Controller, err error) {
	dummy := &txSignin{
		info: newDummyTokenInfo(),
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var content interface{}
	if content, err = dummy.Postcondition(ctx); err != nil {
		return
	}

	subject, ok := content.(*clientDTO.SigninResponseDTO)
	if !ok {
		err = fmt.Errorf("Cannot assert with sign-in response")
		return
	}

	cookie := subject.Cookie
	return session.GetSessionByID(cookie)
}
