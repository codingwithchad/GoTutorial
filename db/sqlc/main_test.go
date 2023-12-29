package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/codingwithchad/util"
	_ "github.com/lib/pq"
)

var TestQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Unable to load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSources)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	TestQueries = New(testDB)
	os.Exit(m.Run())
}
