package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/quandat10/banking/db/sqlc"
	"github.com/quandat10/banking/util"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	// Connect DB failed
	if err != nil {
		log.Fatal("Can not connect db: ", err)
	}

	// Connect DB successful
	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
