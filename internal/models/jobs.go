package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/1-AkM-0/empreGo-web/internal/pagination"
)

type Job struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Link      string `json:"link"`
	Type      string `json:"type"`
	Company   string `json:"company"`
	Source    string `json:"source"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type JobModel struct {
	DB *sql.DB
}

func (jm *JobModel) Insert(job *Job) error {
	stmt := `
	INSERT INTO jobs (title, link,  source, type, company)
	VALUES (?, ?, ?, ?, ?);
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{job.Title, job.Link, job.Source, job.Type, job.Company}

	_, err := jm.DB.ExecContext(ctx, stmt, args...)
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{&job.ID, &job.Title, &job.Link, &job.Source, &job.Type, &job.CreatedAt}

	err := jm.DB.QueryRowContext(ctx, stmt, id).Scan(args...)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (jm *JobModel) GetJobs(filter pagination.Filter) ([]Job, pagination.Metadata, error) {
	stmt := `
	SELECT COUNT(*) OVER(), id, title, link, source, type, created_at, company
	FROM jobs
	ORDER BY created_at DESC
	LIMIT ? OFFSET ?;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := jm.DB.QueryContext(ctx, stmt, filter.Limit(), filter.Offset())
	if err != nil {
		return nil, pagination.Metadata{}, fmt.Errorf("job model getjobs: %w", err)
	}
	defer rows.Close()

	totalRecords := 0
	var jobs []Job

	for rows.Next() {
		var job Job
		args := []any{&totalRecords, &job.ID, &job.Title, &job.Link, &job.Source, &job.Type, &job.CreatedAt, &job.Company}

		if err := rows.Scan(args...); err != nil {
			return nil, pagination.Metadata{}, err
		}
		jobs = append(jobs, job)
	}

	metadata := pagination.CalculateMetada(totalRecords, filter.Page, filter.PageSize)
	return jobs, metadata, nil
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
