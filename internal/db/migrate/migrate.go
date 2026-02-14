package migrate

import (
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/HishamAlkahtani/yt-auto-archiver/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

func MigrateUp() error {
	gormDb, err := db.NewDb()

	if err != nil {
		return fmt.Errorf("failed to get gorm db instance: %w", err)
	}

	sqlDb, err := gormDb.DB()

	if err != nil {
		return fmt.Errorf("failed to get sql instance from gorm db: %w", err)
	}

	dbDriver, err := postgres.WithInstance(sqlDb, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("failed to get postgres db driver: %w", err)
	}

	sourceFS, err := iofs.New(migrationFS, "migrations")

	if err != nil {
		return fmt.Errorf("failed to get migraiton fs: %w", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		sourceFS,
		"postgres",
		dbDriver,
	)

	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	err = m.Up()

	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to run up migrations: %w", err)
		}

		slog.Info("migrated up: no change")
	}

	return nil
}
