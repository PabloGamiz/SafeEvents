package mysql

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func testInit() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Panicf("Got %s while setting up dotenv", err.Error())
	}
}

func TestConnector(t *testing.T) {
	testInit()

	stream, err := OpenStream()
	if err != nil {
		t.Errorf("Got %v error while opening stream", err.Error())
		t.FailNow()
	}

	defer stream.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	query := "SELECT * FROM mastermind.Users"
	var rows *sql.Rows
	if rows, err = stream.QueryContext(ctx, query); err != nil {
		t.Errorf("Got %v error while executing query", err.Error())
		t.FailNow()
	}

	defer rows.Close()

	var id, nickname, email string
	if ok := rows.Next(); !ok {
		t.Errorf("Got no row while getting values from first row")
		t.FailNow()
	}

	if err := rows.Scan(&id, &nickname, &email); err != nil {
		t.Errorf("Got %v error while scanning row's values", err.Error())
		t.FailNow()
	}
}
