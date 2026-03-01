package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

var (
	ErrDuplicateEmail = errors.New("constraint failed: UNIQUE constraint failed: users.email")
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	GithubID  string `json:"github_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (um UserModel) InsertGithub(user *User) error {
	stmt := `
	INSERT INTO users (id, email, username, github_id)
	VALUES (?, ?, ?, ?)
	`

	args := []any{user.ID, user.Email, user.Username, user.GithubID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := um.DB.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return nil
}

func (um UserModel) GetByID(id string) (*User, error) {
	stmt := `
	SELECT id, email, username, github_id 
	FROM users
	WHERE id = ?
	`
	var user User
	args := []any{&user.ID, &user.Email, &user.Username, &user.GithubID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := um.DB.QueryRowContext(ctx, stmt, id).Scan(args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecords
		default:
			return nil, err
		}
	}
	return &user, nil

}

func (um UserModel) GetOrCreateGithubUser(ghUser goth.User) (string, error) {
	var internalID string
	stmt := `
	SELECT id from users
	WHERE github_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := um.DB.QueryRowContext(ctx, stmt, ghUser.UserID).Scan(&internalID)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if err == nil {
		fmt.Printf("retornando id existente: %s", internalID)

		return internalID, nil
	}

	newID := uuid.New().String()
	insertStmt := `
	INSERT INTO users (id, username, email, github_id)
	VALUES (?, ?, ?, ?)
	`
	args := []any{newID, ghUser.NickName, ghUser.Email, ghUser.UserID}
	_, err = um.DB.ExecContext(ctx, insertStmt, args...)
	if err != nil {
		return "", err
	}

	fmt.Printf("retornando novo id: %s", internalID)
	return newID, nil

}
