package data

import (
	"database/sql"
	"errors"
	"strings"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("data conflict")
	ErrDuplicateField = errors.New("duplicate field")
)

func handleSqlDuplicateEntryError(err error) error {
	switch {
	case strings.Contains(err.Error(), "duplicate"):
		return ErrDuplicateField
	default:
		return err
	}

}
func handleSqlNotFoundResultError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrRecordNotFound
	default:
		return err
	}
}

func handleSqlUpdateConflictResultError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrEditConflict
	default:
		return err
	}
}
