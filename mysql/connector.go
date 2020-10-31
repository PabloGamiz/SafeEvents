package mysql

import (
	"database/sql/driver"
	"os"

	"github.com/alvidir/util/pattern/singleton"
)

// Single connInstance of Client
var (
	connInstance = singleton.NewSingleton(initMysqlConn)

	mysqlURL  = os.Getenv(EnvMysqlURL)
	mysqlUser = os.Getenv(EnvMysqlUser)
	mysqlPwd  = os.Getenv(EnvMysqlPwd)
)

// getConnInstance returns the single instance of database.Connector. Multiple calls returns the same instance
func getConnInstance() (conn driver.Connector, err error) {
	var current interface{}
	if current, err = connInstance.GetInstance(); err == nil {
		conn = current.(driver.Connector)
	}

	return
}
