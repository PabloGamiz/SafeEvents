package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"

	"github.com/PabloGamiz/SafeEvents-Backend/api"
	"github.com/alvidir/util/config"
	"github.com/joho/godotenv"
)

const (
	currentEnv = "dev"

	infoSetup = "The server is being started on %s%s"
	infoDone  = "The service has finished successfully"

	errConfigFailed = "Got %s, while setting up service configuration"
	errDotenvConfig = "Service has failed setting up dotenv: %s"
	errListenFailed = "Service has failed listening: %s"
	errServeFailed  = "Service has failed serving: %s"

	envPortKey = "SERVICE_PORT"
	envNetwKey = "SERVICE_NETW"
)

func getMainEnv() ([]string, error) {
	return config.CheckNemptyEnv(
		envPortKey, /*0*/
		envNetwKey /*1*/)
}

func setupDummyUser() error {
	testUser := &client.Client{
		ID:    0,
		Email: "testing@gmail.com",
		Organize: organizer.Organizer{
			ID: 0,
		},
		Assists: assistant.Assistant{
			ID: 0,
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(9000*time.Hour))
	sess, err := session.NewSession(ctx, cancel, testUser)
	if err != nil {
		return err
	}

	log.Printf("Dummy cookie: %v", sess.Cookie())
	return nil
}

func main() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Panicf(errDotenvConfig, err.Error())
	}

	if err := setupDummyUser(); err != nil {
		log.Fatalf("Got %v, while setting up the dummy user", err.Error())
	}

	envs, err := getMainEnv()
	if err != nil {
		log.Fatalf(errConfigFailed, err.Error())
	}

	address := ":" + envs[0]
	log.Printf(infoSetup, envs[1], address)

	lis, err := net.Listen(envs[1], address)
	if err != nil {
		log.Panicf(errListenFailed, err)
	}

	server := api.NewServer()
	if err := http.Serve(lis, server.Router()); err != nil {
		log.Panicf(errServeFailed, err)
	}

	// on finishing
	log.Print(infoDone)
}
