package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database configuration constants
const (
	dbUser     = "your_db_username" // Replace with your database username
	dbPassword = "your_db_password" // Replace with your database password
	dbEndpoint = "your_db_endpoint" // Replace with your Aurora endpoint, e.g., "aurora-instance1.cluster-c1c2c3c4c5.us-west-2.rds.amazonaws.com"
	dbName     = "your_db_name"     // Replace with your database name
)

// ConnectToDatabase creates a connection to the RDS database and returns the *sql.DB object
func ConnectToDatabase() (*sql.DB, error) {
	// Build the database connection string
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbEndpoint, dbName)

	// Establish a connection to the database
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach database: %v", err)
		return nil, err
	}

	log.Println("Connected to database successfully")
	return db, nil
}
