package mysql

import (
	"database/sql"
	"database/sql/driver"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	gormSqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMysqlConn() (interface{}, error) {
	return newMysqlDriver()
}

func newMysqlDriver() (driver driver.Connector, err error) {

	conn := mysql.NewConfig()

	conn.Loc = time.Local
	conn.Timeout = Timeout
	conn.ReadTimeout = Timeout
	conn.WriteTimeout = Timeout

	conn.Addr = os.Getenv("OPEN_NEBULA_IP")
	conn.DBName = os.Getenv("MYSQL_DB")
	conn.User = os.Getenv("MYSQL_USR")
	conn.Passwd = os.Getenv("MYSQL_PWD")
	conn.Net = os.Getenv("SERVICE_NETW")
	conn.ParseTime = true

	return mysql.NewConnector(conn)
}

// OpenStream returns a gateway to the mysql database
func OpenStream() (gormDB *gorm.DB, err error) {
	var conn driver.Connector
	if conn, err = getConnInstance(); err == nil {
		db := sql.OpenDB(conn)
		gormDB, err = gorm.Open(gormSqlDriver.New(gormSqlDriver.Config{
			Conn: db,
		}), &gorm.Config{})
	}

	return
}
