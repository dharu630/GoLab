package Models

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/template"

	_ "github.com/lib/pq" // Import the PostgreSQL driver (or any other driver you need)
)

// Schema structure holds the table definitions
type Schema struct {
	UsersTable string
	PostsTable string
}

// GenerateSQL creates the SQL statements for schema generation
func GenerateSQL() (string, error) {
	// Define a basic schema structure (You can expand this)
	schema := Schema{
		UsersTable: `CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`,
		PostsTable: `CREATE TABLE IF NOT EXISTS posts (
				id SERIAL PRIMARY KEY,
				user_id INT REFERENCES users(id),
				content TEXT NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`,
	}

	// SQL template for migration
	sqlTemplate := `
		-- Up Migration
		{{.UsersTable}}

		{{.PostsTable}}
	`

	// Create the template
	tmpl, err := template.New("schema").Parse(sqlTemplate)
	if err != nil {
		return "", err
	}

	// Use a bytes.Buffer to capture the output of the template execution
	var sqlScriptBuffer bytes.Buffer
	err = tmpl.Execute(&sqlScriptBuffer, schema)
	if err != nil {
		return "", err
	}

	// Return the generated SQL script from the buffer
	return sqlScriptBuffer.String(), nil
}

// WriteMigrationFiles writes the SQL to migration files
func WriteMigrationFiles(upSQL, migrationPath string) error {
	// Generate file names
	upFile := fmt.Sprintf("%s.up.sql", migrationPath)
	downFile := fmt.Sprintf("%s.down.sql", migrationPath)

	// Write the up migration file
	err := os.WriteFile(upFile, []byte(upSQL), 0644)
	if err != nil {
		return err
	}

	// Define downSQL for rolling back the migration
	downSQL := `
		DROP TABLE IF EXISTS posts;
		DROP TABLE IF EXISTS users;
	`

	// Write the down migration file (reverse operations)
	err = os.WriteFile(downFile, []byte(downSQL), 0644)
	if err != nil {
		return err
	}

	log.Printf("Migration files created: %s, %s", upFile, downFile)
	return nil
}

// RunMigrations applies the migrations to the database
func RunMigrations(db *sql.DB, migrationPath string) error {
	// Read the "up" migration SQL from file
	upFile := fmt.Sprintf("%s.up.sql", migrationPath)
	upSQL, err := os.ReadFile(upFile)
	if err != nil {
		return fmt.Errorf("failed to read up migration file: %v", err)
	}

	// Execute the SQL for the migration
	_, err = db.Exec(string(upSQL))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %v", err)
	}

	log.Printf("Migration applied successfully from %s", upFile)
	return nil
}
