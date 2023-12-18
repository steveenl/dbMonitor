package main

import (
	"log"
)

func main() {
	// Establish a connection to the database
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Call the data collection function
	CollectData()
}
