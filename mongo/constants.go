package mongo

import "time"

const (
	// EnvMongoAPI represents the environment variable where the api key is located
	EnvMongoAPI = "MONGO_API"

	// Database consultable by the app
	Database = "safe-events"
	// Timeout for any database request
	Timeout = 10 * time.Second
)
