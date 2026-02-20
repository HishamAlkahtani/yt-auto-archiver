package main

import (
	"log/slog"
	"os"

	"github.com/HishamAlkahtani/yt-auto-archiver/internal/migrate"
)

func main() {
	err := migrate.MigrateUp()

	if err != nil {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}
}
