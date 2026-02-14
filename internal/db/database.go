package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HishamAlkahtani/yt-auto-archiver/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

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

	return &DB{DB: sqlDb}, nil
}

func (db *DB) UpdateVideo(vid *model.Video) error {
	return nil
}

func (db *DB) GetVideosByChannelAndStatus(channelId string, status model.VideoStatus) ([]model.Video, error) {
	return nil, nil
}

func (db *DB) GetVideosByChannel(channelId string) ([]model.Video, error) {
	return nil, nil
}

func (db *DB) GetVideo(videoId string) (*model.Video, error) {
	return nil, nil
}

func (db *DB) UpdateChannel(channel *model.Channel) error {
	return nil
}

func (db *DB) GetChannel(channelId string) (*model.Channel, error) {
	return nil, nil
}
