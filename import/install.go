package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection string using environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbHost == "" || dbUser == "" || dbPassword == "" {
		log.Fatal("Missing required environment variables: DB_HOST, DB_USER, DB_PASSWORD")
	}

	// Set defaults if not provided
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbName == "" {
		dbName = "defaultdb"
	}
	if dbSSLMode == "" {
		dbSSLMode = "require"
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Read the SQL file
	sqlFilePath := "database-dump.sql" // Adjust this to your .sql file path
	sqlContent, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}

	// Split the SQL content into individual statements
	// This is a basic split; it assumes statements end with semicolons
	statements := strings.Split(string(sqlContent), ";\n")

	// Execute each statement
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue // Skip empty statements
		}

		// Remove comment lines from the statement
		lines := strings.Split(stmt, "\n")
		var cleanedLines []string
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "--") {
				continue // Skip comment lines
			}
			cleanedLines = append(cleanedLines, trimmed)
		}

		cleanedStmt := strings.Join(cleanedLines, " ")
		if cleanedStmt == "" {
			continue // Skip if the statement is empty after cleaning
		}

		// Execute the cleaned statement
		_, err := db.Exec(cleanedStmt)
		if err != nil {
			log.Printf("Failed to execute statement: %v\nStatement: %s\n", err, cleanedStmt)
			// Continue with next statement even if one fails
			return
		}
		fmt.Printf("Executed: %s\n", cleanedStmt[:min(50, len(cleanedStmt))]+"...") // Log first 50 chars
	}

	fmt.Println("SQL script execution completed.")
}

// Helper function to get min of two ints
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
