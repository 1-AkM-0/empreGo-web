package storage

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"io/fs"
	_ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {

	db, err := sql.Open("sqlite", "vagas.db?_pragma=journal_mode(WAL)&_busy_timeout=5000")
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	fmt.Println("Conectado ao database")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("sqlite3")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}
	return nil
}
