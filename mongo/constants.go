package mongo

import "time"

const (
	// EnvMongoURL represents the environment variable where the mongo's url key is located
	EnvMongoURL = "MONGO_URL"

	// Database consultable by the app
	Database = "safe-events-db"
	// Timeout for any database request
	Timeout = 10 * time.Second
)
