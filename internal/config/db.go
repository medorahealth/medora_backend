package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// ConnectDatabase loads environment variables and establishes a connection to the database.
// It returns the database connection object or an error.
func ConnectDatabase() (*sql.DB, error) {
	// Attempt to load the .env file from the root of the project.
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Relying on OS environment variables.")
	}

	// Read the DATABASE_URL from the environment.
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("FATAL: DATABASE_URL environment variable is not set.")
	}

	// Open a connection to the database.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection is alive.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
