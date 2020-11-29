package client

import (
	"context"
	"testing"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/joho/godotenv"
)

func newSessionInfo(cookie string) clientDTO.LogoutRequestDTO {
	return clientDTO.LogoutRequestDTO{
		Cookie: cookie,
	}
}

func TestLogout_postcondition(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Got error %s; while loading dotenv", err.Error())
	}

	signin := &txSignin{
		info: newTestTokenInfo(),
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	response, err := signin.Postcondition(ctx)
	if err != nil {
		t.Fatalf("Got error %s; while signing in dummy user", err.Error())
	}

	resp, ok := response.(*clientDTO.SigninResponseDTO)
	if !ok {
		t.Fatalf("Got an unexpected response type, want *SigninResponseDTO")
	}

	subject := &txLogout{
		request: newSessionInfo(resp.Cookie),
	}

	if response, err = subject.Postcondition(ctx); err != nil {
		t.Fatalf("Got error %s; while executing logout Postcondition", err.Error())
	}

	result, ok := response.(*clientDTO.LogoutResponseDTO)
	if !ok {
		t.Fatalf("Got an unexpected response type, want *LogoutResponseDTO")
	}

	if result.Cookie != resp.Cookie {
		t.Errorf("Got unexpected cookie %v as response, want %v", result.Cookie, resp.Cookie)
	}

	if now := time.Now().Unix(); result.Deadline > now {
		t.Errorf("Got unexpected deadline %v, should be lower than now %v", result.Deadline, now)
	}
}
