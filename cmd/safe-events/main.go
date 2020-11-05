package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PabloGamiz/SafeEvents-Backend/api"
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
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

	// Migració de structs del Model (Es fa automatica si tenen els tags ben definits).
	// db.AutoMigrate(&service.Service{})

	// Afegir files a les taules de la BBDD. Em suposo que se li pot passar l'struct del model ja construit, no cal construir-lo "in situ".
	db.Create(&service.Service{
		Name:        "service test",
		Description: "description of service test",
		Kind:        1,
		Location: location.Location{
			Name:        "location test",
			Address:     "address test",
			Coordinates: "101010",
			Extension:   10},
		Products: []product.Product{{
			Name:        "product test",
			Description: "description of product test",
			Price:       10,
			Status:      1}}})

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
