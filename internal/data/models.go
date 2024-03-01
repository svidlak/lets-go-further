package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

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
