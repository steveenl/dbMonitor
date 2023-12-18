package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbUser     string
	dbPassword string
	dbEndpoint string
	dbName     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbEndpoint = os.Getenv("DB_ENDPOINT")
	dbName = os.Getenv("DB_NAME")
}

func main() {
	// Establish a connection to the database
	db, err := ConnectToDatabase(false)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	CollectData()

	CheckForAnomalies(db)
}
