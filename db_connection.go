package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase(isLocal bool) (*sql.DB, error) {
	if isLocal {
		// Return a mock database connection for local testing
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			return nil, err
		}
		mock.ExpectPing()
		return db, nil
	} else {
		// Connect to AWS Aurora for production
		dbUser := os.Getenv("AWS_DB_USER")
		dbPassword := os.Getenv("AWS_DB_PASSWORD")
		dbEndpoint := os.Getenv("AWS_DB_ENDPOINT")
		dbName := os.Getenv("AWS_DB_NAME")

		dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbEndpoint, dbName)
		db, err := sql.Open("mysql", dbConnectionString)
		if err != nil {
			log.Printf("Unable to connect to database: %v", err)
			return nil, err
		}

		if err = db.Ping(); err != nil {
			log.Printf("Unable to reach database: %v", err)
			return nil, err
		}

		log.Println("Connected to database successfully")
		return db, nil
	}
}
