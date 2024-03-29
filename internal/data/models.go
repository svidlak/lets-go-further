package data

import (
	"database/sql"
)

type Models struct {
	Movies      MovieModel
	Users       UserModel
	Permissions PermissionModel
	Tokens      TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}
