package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// LogEntry represents a parsed log entry
type LogEntry struct {
	Timestamp string
	LogLevel  string
	Message   string
}

// ParseLogEntry parses a single log entry and returns a LogEntry struct
func ParseLogEntry(logLine string) (*LogEntry, error) {
	// Adjust the regex according to your log format
	logPattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] (.+)`)
	matches := logPattern.FindStringSubmatch(logLine)

	if len(matches) < 4 {
		return nil, fmt.Errorf("invalid log entry")
	}

	return &LogEntry{
		Timestamp: matches[1],
		LogLevel:  matches[2],
		Message:   matches[3],
	}, nil
}

// AnalyzeLogs processes the log file for analysis
func AnalyzeLogs(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logEntry, err := ParseLogEntry(scanner.Text())
		if err != nil {
			fmt.Println("Error parsing log entry:", err)
			continue
		}

		// Perform your analysis here (e.g., look for errors, warnings, etc.)
		fmt.Printf("Parsed Log Entry: %+v\n", logEntry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading log entries:", err)
	}
}
