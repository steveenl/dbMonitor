package main

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockEmailSender struct{}

func (m MockEmailSender) SendEmail(recipient string, subject string, body string) error {
	// Mock implementation for testing
	return nil
}

func TestAnalyzeData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"query", "execution_time"}).
		AddRow("SELECT * FROM table", 1100).
		AddRow("UPDATE table SET column = value", 1200)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT query, execution_time FROM queries WHERE execution_time > ?;`)).
		WithArgs(1000).
		WillReturnRows(rows)

	AnalyzeData(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestCheckForAnomalies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"query", "execution_time"}).
		AddRow("SELECT * FROM table", 5100).
		AddRow("UPDATE table SET column = value", 5200)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT query, execution_time FROM queries WHERE execution_time > ?;`)).
		WithArgs(5000).
		WillReturnRows(rows)

	CheckForAnomalies(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
