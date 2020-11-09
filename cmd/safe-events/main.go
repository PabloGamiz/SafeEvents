package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PabloGamiz/SafeEvents-Backend/api"
	"github.com/joho/godotenv"
)

const (
	currentEnv = "dev"

	infoSetup = "The server is being started on %s%s"
	infoDone  = "The service has finished successfully"

	errConfigFailed = "Got %s, while setting up service configuration"
	errDotenvConfig = "Service has failed setting up dotenv: %s"
	errListenFailed = "Service has failed listening: %v"
	errServeFailed  = "Service has failed serving: %v"

	envPortKey = "SERVICE_PORT"
	envNetwKey = "SERVICE_NETW"

	defaultPort    = "9090"
	defaultNetwork = "tcp"
)

func network() string {
	if value, ok := os.LookupEnv(envNetwKey); ok {
		return value
	}

	return defaultNetwork
}

func address() (address string) {
	address = defaultPort
	if value, ok := os.LookupEnv(envPortKey); ok {
		address = value
	}

	if address[0] != ':' {
		address = fmt.Sprintf(":%s", address)
	}

	return
}

func main() {
	address := address()
	network := network()
	log.Printf(infoSetup, network, address)

	env := godotenv.Load(".env." + currentEnv)
	if env != nil {
		log.Fatalf("Error loading .env.%s file", currentEnv)
	}

	lis, err := net.Listen(network, address)
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
