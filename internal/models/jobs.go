package models

import (
	"database/sql"
	"fmt"
)

type Job struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Link      string `json:"link"`
	Type      string `json:"type"`
	Source    string `json:"source"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type JobModel struct {
	DB *sql.DB
}

func (jm *JobModel) Insert(job *Job) error {
	stmt := `
	INSERT INTO jobs (title, link,  source, type)
	VALUES (?, ?, ?, ?);
	`

	args := []any{job.Title, job.Link, job.Source, job.Type}

	_, err := jm.DB.Exec(stmt, args...)
	if err != nil {
		return fmt.Errorf("job model insert: %w", err)
	}
	return nil
}

func (jm *JobModel) GetJobByID(id int) (*Job, error) {
	stmt := `
	SELECT id, title, link, source, type, created_at
	FROM jobs
	WHERE id = ?
	`
	var job Job

	args := []any{&job.ID, &job.Title, &job.Link, &job.Source, &job.Type, &job.CreatedAt}

	err := jm.DB.QueryRow(stmt, id).Scan(args...)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (jm *JobModel) GetJobs() ([]Job, error) {
	stmt := `
	SELECT id, title, link, source, type, created_at
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
		args := []any{&job.ID, &job.Title, &job.Link, &job.Source, &job.Type, &job.CreatedAt}

		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (jm *JobModel) Exists(link string) bool {
	var exists int
	stmt := `
	SELECT 1 from jobs
	WHERE link = ?
	LIMIT 1
	`
	err := jm.DB.QueryRow(stmt, link).Scan(&exists)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		fmt.Println("JobModel exists:", err)
		return true
	}
	return true

}
