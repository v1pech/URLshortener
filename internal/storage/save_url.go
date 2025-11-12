package storage

import "fmt"

func (s *Storage) SaveURL(url, alias string) (int8, error) {
	const op = "storage.SaveURL"
	var exists = false
	_ = s.Database.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE alias = ?)", alias).Scan(&exists)
	if exists {
		return ErrAliasExists, fmt.Errorf("%s: %s", op, "alias already exists")
	}
	_, err := s.Database.Exec("INSERT INTO urls (alias, url) VALUES (?, ?)", alias, url)
	if err != nil {
		return ErrDatabase, fmt.Errorf("%s: %w", op, err)
	}
	return NoError, nil
}
