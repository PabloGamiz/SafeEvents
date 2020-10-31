package mysql

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/alvidir/util/config"
	"github.com/go-sql-driver/mysql"
)

func initMysqlConn() (interface{}, error) {
	return newMysqlDriver()
}

func newMysqlDriver() (driver driver.Connector, err error) {
	var envs []string
	if envs, err = config.CheckNemptyEnv(EnvMysqlURL, EnvMysqlUser, EnvMysqlPwd, EnvMysqlDB); err != nil {
		return
	}

	conn := mysql.NewConfig()

	conn.Loc = time.Local
	conn.Timeout = Timeout
	conn.ReadTimeout = Timeout
	conn.WriteTimeout = Timeout

	conn.Addr = envs[0]
	conn.DBName = envs[3]
	conn.Passwd = envs[2]
	conn.User = envs[1]
	conn.Net = "tcp"

	return mysql.NewConnector(conn)
}

// OpenStream returns a gateway to the mysql database
func OpenStream() (db *sql.DB, err error) {
	var conn driver.Connector
	if conn, err = getConnInstance(); err == nil {
		db = sql.OpenDB(conn)
	}

	return
}
