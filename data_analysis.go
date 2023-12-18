package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/smtp"
)

// AnalyzeData processes the collected data to identify slow-running queries and bottlenecks
func AnalyzeData(db *sql.DB) {
	query := `SELECT query, execution_time FROM queries WHERE execution_time > ?;`
	threshold := 1000 // Threshold in milliseconds

	rows, err := db.Query(query, threshold)
	if err != nil {
		fmt.Println("Error executing analysis query:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Slow Queries:")
	for rows.Next() {
		var query string
		var executionTime int
		err := rows.Scan(&query, &executionTime)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("Query: %s, Execution Time: %dms\n", query, executionTime)
	}
}

func CheckForAnomalies(db *sql.DB) (bool, error) {
	query := `SELECT query, execution_time FROM queries WHERE execution_time > ?;`
	threshold := 5000 // Threshold in milliseconds for unusually long queries

	rows, err := db.Query(query, threshold)
	if err != nil {
		fmt.Println("Error executing anomaly detection query:", err)
		return false, err
	}
	defer rows.Close()

	anomalyDetected := false
	for rows.Next() {
		var query string
		var executionTime int
		if err := rows.Scan(&query, &executionTime); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		// Anomaly detected
		anomalyDetected = true
		fmt.Printf("Slow query detected: %s, Execution Time: %dms\n", query, executionTime)
	}

	return anomalyDetected, nil
}

// Send email alerts
func sendAlert(message string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "your-email@gmail.com", "your-password", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.com"}
	msg := []byte("To: recipient@example.com\r\n" +
		"Subject: Database Alert!\r\n" +
		"\r\n" +
		message + "\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "your-email@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
