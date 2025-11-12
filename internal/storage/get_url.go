package storage

import (
	"database/sql"
	"fmt"
)

func (db *Storage) GetUrl(alias string) (string, int8, error) {
	const op = "storage.GetUrl"

	var url string

	err := db.Database.QueryRow("SELECT url FROM urls WHERE alias = ?", alias).Scan(&url)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrAliasNotExists, fmt.Errorf("%s: %s", op, "url not found")
		} else {
			return "", ErrDatabase, fmt.Errorf("%s: %w", op, err)
		}
	}

	return url, NoError, nil
}
