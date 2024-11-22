package main

import (
	"database/sql"
	"fmt"
	"go_migrate/Models"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=taskuser password=12345678 dbname=taskmanager sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	upSQL, err := Models.GenerateSQL()
	if err != nil {
		log.Fatalf("Failed to generate SQL schema: %v", err)
	}

	migrationPath := "migrations/6_create_schema"

	err = Models.WriteMigrationFiles(upSQL, migrationPath)
	if err != nil {
		log.Fatalf("Failed to write migration files: %v", err)
	}

	err = Models.RunMigrations(db, migrationPath)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations completed successfully!")
}
