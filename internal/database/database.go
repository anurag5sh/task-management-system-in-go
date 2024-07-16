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

type Database struct {
	Ctx context.Context
	Db  *sql.DB
}

func NewDatabase() (*Database, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatalf("Unable to open database: %v", err)
		return nil, err
	}

	if _, err = db.ExecContext(ctx, ddl); err != nil {
		log.Fatalf("Unable to create database: %v", err)
		return nil, err
	}

	return &Database{Db: db, Ctx: ctx}, nil
}
