package storage

import "database/sql"

type Storage struct {
	Database *sql.DB
}

const (
	NoError           = 0
	ErrDatabase       = -1
	ErrAliasExists    = 1
	ErrAliasNotExists = 2
)
