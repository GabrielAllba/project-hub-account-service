package dbutil

import (
	"database/sql"
	"fmt"
	"os"
)

func RunMigrationsFromFile(db *sql.DB, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}
