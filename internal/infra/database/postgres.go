package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type psql struct{
	Db *sql.DB
}

func New(hostname, port, user, password, dbName string) *psql {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostname, port, user, password, dbName)

	db, err := sql.Open("postgres", dsn)

	if (err != nil) {
		log.Fatalf("error connecting to the database: %v", err)
	}

	return &psql{
		Db: db,
	}
}