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

type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
		GetAll(title string, genres []string, filter Filter) ([]*Movie, Metadata, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
