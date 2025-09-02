package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed migrations/*.sql
var migrationFs embed.FS

type DB struct {
	*sql.DB
}

func NewDB(dbPath string) (*DB, error) {
	dir := filepath.Dir(dbPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create db directory: %w", err)
	}

	sqlDb, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := sqlDb.Ping(); err != nil {
		sqlDb.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := runMigrations(sqlDb); err != nil {
		sqlDb.Close()
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	return &DB{DB: sqlDb}, nil
}

func runMigrations(db *sql.DB) error {
	// // dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	// if err != nil {
	// 	return fmt.Errorf("failed to create database driver: %w", err)
	// }

	// migrator, err := migrate.NewWithDatabaseInstance()

	// if err != nil {
	// 	return fmt.Errorf("failed to create migrator: %w", err)
	// }
	// defer migrator.Close()

	// if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
	// 	return fmt.Errorf("failed to run migrations: %W", err)
	// }

	// return nil
}
