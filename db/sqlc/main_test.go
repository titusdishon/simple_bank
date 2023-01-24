package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/titusdishon/simple_bank/util"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..", "test")
	if err != nil {
		log.Fatal(" Cannot load config: ", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(" Cannot connect to db: ", err)
	}
	testQueries = New(testDb)
	os.Exit(m.Run())
}
