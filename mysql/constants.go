package mysql

import "time"

const (
	// EnvMysqlURL represents the environment variable where the mysql url is located
	EnvMysqlURL = "MYSQL_HOST"
	// EnvMysqlUser represents the environment variable where the mysql username is located
	EnvMysqlUser = "MYSQL_USR"
	// EnvMysqlPwd represents the environment variable where the mysql password is located
	EnvMysqlPwd = "MYSQL_PWD"
	// EnvMysqlDB represents the environment variable where the mysql database is located
	EnvMysqlDB = "MYSQL_DB"

	// Timeout for any database request
	Timeout = 10 * time.Second
)
