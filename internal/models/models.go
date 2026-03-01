package models

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecords = errors.New("recurso não encontrado")
)

type Models struct {
	UserModel        UserModel
	JobModel         JobModel
	ApplicationModel ApplicationModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		JobModel: JobModel{
			DB: db,
		},
		UserModel: UserModel{
			DB: db,
		},
		ApplicationModel: ApplicationModel{
			DB: db,
		},
	}
}
