package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func SetupStorage(dbPath string) (*Storage, error) {
	op := "storage.SetupStorage"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls
	(
	id INTEGER PRIMARY KEY,
	alias TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db}, nil
}
