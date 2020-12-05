package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/api"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql/migration"
	clientTX "github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
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

func setup() (err error) {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err = godotenv.Load(); err != nil {
		return
	}

	if err = migration.MigrateTables(); err != nil {
		return
	}

	if err = clientTX.SetupDummyUser(); err != nil {
		return
	}

	return
}

func main() {
	fmt.Println(time.Now())
	if err := setup(); err != nil {
		log.Fatalf(err.Error())
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
