package models

import "database/sql"

type Models struct {
	JobModel JobModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		JobModel: JobModel{
			DB: db,
		},
	}
}
