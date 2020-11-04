package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PabloGamiz/SafeEvents-Backend/api"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	mysql "github.com/PabloGamiz/SafeEvents-Backend/mysql"
)

const (
	currentEnv = "dev"

	infoSetup = "The server is being started on %s%s"
	infoDone  = "The service has finished successfully"

	errDotenvConfig = "Service has failed setting up dotenv: %s"
	errListenFailed = "Service has failed listening: %s"
	errServeFailed  = "Service has failed serving: %s"

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

func test() {
	db, err := mysql.OpenStream()
	if err != nil {
		log.Printf("Got %v error while opening stream", err.Error())
		return
	}

	db.AutoMigrate(&event.Event{})
}

func main() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	address := address()
	network := network()
	log.Printf(infoSetup, network, address)

	test()

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
