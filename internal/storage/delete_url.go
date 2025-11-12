package storage

import (
	"database/sql"
	"fmt"
)

func (db *Storage) DeleteURL(alias string) (int8, error) {
	const op = "storage.DeleteURL"

	_, err := db.Database.Exec("DELETE FROM urls WHERE alias = ?", alias)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrAliasNotExists, fmt.Errorf("%s: %s", op, "alias not found")
		} else {
			return ErrDatabase, fmt.Errorf("%s: %w", op, err)
		}
	}

	return NoError, nil
}
