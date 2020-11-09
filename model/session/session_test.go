package session

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

func getMapLength(m *sync.Map) int {
	count := 0
	m.Range(func(key interface{}, sess interface{}) bool {
		log.Printf("A session exists under key %s", key)
		count++
		return true
	})

	return count
}

func TestGetSessionByEmail(t *testing.T) {
	client := &client.Client{
		Email: "testing@gmail.com",
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	subject, err := NewSession(ctx, cancel, client)
	if err != nil {
		t.Fatalf("Got error %s, while getting new session", err.Error())
	}

	want := 1
	if got := getMapLength(AllInstancesByID); got != want {
		t.Errorf("Got len(AllInstancesByID) = %v, want %v", got, want)
	}

	if got := getMapLength(AllInstancesByEmail); got != want {
		t.Errorf("Got len(AllInstancesByEmail) = %v, want %v", got, want)
	}

	sess, err := GetSessionByEmail(subject.GetEmail())
	if err != nil {
		t.Fatalf("Got error %s, while getting session by email %s", err.Error(), subject.GetEmail())
	}

	if sess.Cookie() != subject.Cookie() {
		t.Errorf("Got session cookie %s, want %s", sess.Cookie(), subject.Cookie())
	}
}
