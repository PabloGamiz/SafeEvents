package mongo

import (
	"os"

	"github.com/alvidir/util/pattern/singleton"
	"go.mongodb.org/mongo-driver/mongo"
)

// Single clientInstance of Client
var (
	clientInstance = singleton.NewSingleton(initMongoClient)
	mongoAPI       = os.Getenv(EnvMongoAPI)
)

// GetClientInstance returns the single instance of Client. Multiple calls returns the same instance
func GetClientInstance() (client *mongo.Client, err error) {
	var current interface{}
	if current, err = clientInstance.GetInstance(); err == nil {
		client = current.(*mongo.Client)
	}

	return
}
