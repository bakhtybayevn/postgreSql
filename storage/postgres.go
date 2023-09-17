package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config represents the configuration for the database connection.
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

// NewConnection creates a new database connection using the provided configuration.
func NewConnection(config *Config) (*gorm.DB, error) {
	// Construct the Data Source Name (DSN) for the PostgreSQL database.
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	// Open a connection to the database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
