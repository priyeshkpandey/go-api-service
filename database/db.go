package database

import (
	"database/sql"
	"log"

	"api.service.go/go-api-service/common"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(dbFilePath string) {
	db, err := sql.Open("sqlite3", dbFilePath)
	if !common.HasError(err) {
		log.Printf("Successfully opened DB: %v", db)
	} else {
		log.Fatalf("Error opening DB: %v", err)
	}
}
