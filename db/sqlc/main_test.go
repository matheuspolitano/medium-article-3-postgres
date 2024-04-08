package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jpfuentes2/go-env/autoload"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("cannot connet to db:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
