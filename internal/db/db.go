// internal/db/db.go
package db

import (
	"database/sql"
	"fmt"	

	_ "github.com/lib/pq"
	"github.com/rezalaal/coral/config"
)

func Connect() (*sql.DB, error) {
	cfg, err := config.Load()
	connStr := cfg.DatabaseURL
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
