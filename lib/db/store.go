package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/litmus-zhang/task-manager/lib/config"
)

type Store interface {
	Querier
}
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(config *config.Config) Store {
	log.Printf("Connecting to %s, %s", config.DB.DbSource, config.DB.DbDriver)

	db, err := sql.Open(config.DB.DbDriver, config.DB.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
