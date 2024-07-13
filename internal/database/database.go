package database

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed sqlc/schema.sql
var ddl string

type Database interface {
	NewDatabase()
}

type SqliteDatabase struct {
	ctx context.Context
	db  *sql.DB
}

func (sqlite *SqliteDatabase) NewDatabase() (*sql.DB, error) {
	ctx := context.Background()

	var err error
	sqlite.db, err = sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatalf("Unable to open database: %v", err)
		return nil, err
	}

	if _, err = sqlite.db.ExecContext(ctx, ddl); err != nil {
		log.Fatalf("Unable to create database: %v", err)
		return nil, err
	}

	return sqlite.db, nil
}
