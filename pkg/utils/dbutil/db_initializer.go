package dbutil

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func EnsureDatabaseExists(driverName, dbName, defaultDBSource string) error {
	defaultDB, err := sql.Open(driverName, defaultDBSource)
	if err != nil {
		return fmt.Errorf("failed to connect to default DB: %w", err)
	}
	defer defaultDB.Close()

	_, err = defaultDB.Exec("CREATE DATABASE \"" + dbName + "\"")
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P04" {
			log.Printf("Database %s already exists", dbName)
			return nil
		}
		return fmt.Errorf("error creating database: %w", err)
	}

	log.Printf("Database %s created successfully", dbName)
	return nil
}
