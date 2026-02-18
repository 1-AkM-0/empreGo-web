package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Job struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Status    string    `json:"status"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JobModel struct {
	DB *sql.DB
}

func (jm *JobModel) Insert(job *Job) error {
	stmt := `
	INSERT INTO jobs (title, link,  source)
	VALUES (?, ?, ?);
	`

	args := []any{job.Title, job.Link, job.Source}

	_, err := jm.DB.Exec(stmt, args...)
	if err != nil {
		return fmt.Errorf("job model insert: %w", err)
	}
	return nil
}

func (jm *JobModel) GetJobByID(id int) (*Job, error) {
	stmt := `
	SELECT id, title, link, source
	FROM jobs
	WHERE id = ?
	`
	var job Job

	args := []any{&job.ID, &job.Title, &job.Link, &job.Source}

	err := jm.DB.QueryRow(stmt, id).Scan(args...)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (jm *JobModel) GetJobs() ([]Job, error) {
	stmt := `
	SELECT id, title, link, source
	FROM jobs;
	`

	rows, err := jm.DB.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("job model getjobs: %w", err)
	}
	defer rows.Close()

	var jobs []Job

	for rows.Next() {
		var job Job
		args := []any{&job.ID, &job.Title, &job.Link, &job.Source}

		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
