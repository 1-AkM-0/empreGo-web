package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("constraint failed: UNIQUE constraint failed: users.email")
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	GithubID  string    `json:"github_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (um UserModel) UpsertGithub(user *User) error {
	stmt := `
	INSERT INTO users (id, email, username, github_id)
	VALUES (?, ?, ?, ?)
	ON CONFLICT(github_id) DO NOTHING
	RETURNING id; 
	`

	args := []any{user.ID, user.Email, user.Username, user.GithubID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := um.DB.QueryRowContext(ctx, stmt, args...).Scan(&user.ID)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.email") {
			return ErrDuplicateEmail
		}
		return err
	}
	return nil
}
