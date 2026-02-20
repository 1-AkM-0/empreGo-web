package models

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecords = errors.New("recurso não encontrado")
)

type Models struct {
	UserModel UserModel
	JobModel  JobModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		JobModel: JobModel{
			DB: db,
		},
		UserModel: UserModel{
			DB: db,
		},
	}
}
