package main

import (
	"testing"
)

func TestConnectToDatabase(t *testing.T) {
	db, err := ConnectToDatabase(true) // Pass true to use the local (mock) database
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}

	if db == nil {
		t.Errorf("Expected a non-nil database object, got nil")
	}
}
