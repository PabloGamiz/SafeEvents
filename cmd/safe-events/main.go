package main

import (
	"log"
	"net"
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api"
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	mysql "github.com/PabloGamiz/SafeEvents-Backend/mysql"
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
	if err := godotenv.Load(); err != nil {
		log.Panicf(errDotenvConfig, err.Error())
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
