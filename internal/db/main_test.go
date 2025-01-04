package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/litmus-zhang/task-manager/internal/config"
)

var testQueries *Queries
var testDB *sql.DB

const (
	VAULT_ADDR  = "http://127.0.0.1:8200"
	VAULT_TOKEN = "root"
	VAULT_PATH  = "secret/data/task-manager"
)

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("cannot setup a new config:", err)
	}

	testDB, err = sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
	testDB.Close()
}
