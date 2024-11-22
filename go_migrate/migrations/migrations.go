package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection string
	dbURL := "postgres://taskuser:12345678@localhost:5432/taskmanager?sslmode=disable"

	// Location of migration files
	migrationsDir := "file://C:/Users/ndhar/Documents/go_migrate/migrations"

	// Open a connection to the database
	db, err := openDatabase(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create a database instance for migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create database instance: %v", err)
	}

	// Create a migrate instance
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, "postgres", driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	// Perform migrations
	if len(os.Args) < 2 {
		log.Fatalf("Please provide a command: up, down, or status")
	}

	command := os.Args[1]
	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migrations applied successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migrations rolled back successfully.")
	case "status":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("Could not fetch migration status: %v", err)
		}
		fmt.Printf("Current Version: %d, Dirty: %t\n", version, dirty)
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

// openDatabase opens a database connection
func openDatabase(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	return db, nil
}
