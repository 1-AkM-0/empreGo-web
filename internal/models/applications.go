package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID        string `json:"id"`
	UserID    string `json:"-"`
	JobID     int    `json:"job_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ApplicationResponse struct {
	ApplicationID        string `json:"application_id"`
	JobTitle             string `json:"job_title"`
	JobLink              string `json:"job_link"`
	ApplicationStatus    string `json:"status"`
	JobType              string `json:"job_type"`
	JobSource            string `json:"job_source"`
	ApplicationCreatedAt string `json:"application_created_at"`
}

type ApplicationModel struct {
	DB *sql.DB
}

func (am ApplicationModel) Insert(application *Application) error {
	id := uuid.New().String()
	stmt := `
	INSERT INTO applications (id, user_id, job_id)
	VALUES (?, ?, ?)
	RETURNING id, created_at
	`

	args := []any{id, application.UserID, application.JobID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return am.DB.QueryRowContext(ctx, stmt, args...).Scan(&application.ID, &application.CreatedAt)
}

func (am ApplicationModel) GetAll(userID string) ([]ApplicationResponse, error) {
	stmt := `
	SELECT a.id, j.title, j.link, j.source, a.created_at, j.type, a.status
  FROM applications AS a 
  JOIN jobs AS j
  ON j.id = a.job_id
	WHERE a.user_id = ?;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	applications := []ApplicationResponse{}

	rows, err := am.DB.QueryContext(ctx, stmt, userID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecords
		default:
			return nil, err
		}
	}

	for rows.Next() {
		var applicationResponse ApplicationResponse

		err := rows.Scan(&applicationResponse.ApplicationID,
			&applicationResponse.JobTitle,
			&applicationResponse.JobLink,
			&applicationResponse.JobSource,
			&applicationResponse.ApplicationCreatedAt,
			&applicationResponse.JobType,
			&applicationResponse.ApplicationStatus,
		)
		if err != nil {
			return nil, err
		}
		applications = append(applications, applicationResponse)
	}

	return applications, nil
}

// func (am ApplicationModel) Delete(applicationID string) error
func (am ApplicationModel) Update(userID, status, applicationID string) error {
	stmt := `
	UPDATE applications
	SET status = ?
	WHERE id = ? AND user_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := am.DB.ExecContext(ctx, stmt, status, applicationID, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return ErrNoRecords
	}

	if err != nil {
		return err
	}
	return nil
}

//func (am ApplicationModel) GetApplicationByID(applicationID string) (*Application, error)
