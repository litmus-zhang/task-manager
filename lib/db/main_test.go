package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/litmus-zhang/task-manager/lib/config"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	store := NewStore(cfg)
	testDB = store.(*SQLStore).db
	testQueries = New(testDB)

	testDB.Close()
	os.Exit(m.Run())
}
