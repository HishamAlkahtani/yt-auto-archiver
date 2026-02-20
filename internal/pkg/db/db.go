package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbHost = "postgres_db"
const dbPort = 5432

func NewDb() (*gorm.DB, error) {
	config, err := newConfig()

	if err != nil {
		return nil, fmt.Errorf("failed to load db config: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		dbHost,
		config.User,
		config.Password,
		config.DB,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
