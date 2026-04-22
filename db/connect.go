package db

import (
	"database/sql"
	"log"
	"sync"
)

var (
	dbInstance *sql.DB
	dbErr      error
	once       sync.Once
)

func ReturnDb() *sql.DB {
	once.Do(func() {
		dbInstance, dbErr = Connect()
	})

	if dbErr != nil {
		log.Fatalf("Falha na conexão: %v", dbErr)
	}

	return dbInstance
}
