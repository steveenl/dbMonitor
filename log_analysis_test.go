package main

import (
    "testing"
)

func TestParseLogEntry(t *testing.T) {
    sampleLog := "2021-01-02 15:04:05 [INFO] Sample log message"
    expected := &LogEntry{
        Timestamp: "2021-01-02 15:04:05",
        LogLevel:  "INFO",
        Message:   "Sample log message",
    }

    result, err := ParseLogEntry(sampleLog)
    if err != nil {
        t.Errorf("Error parsing log entry: %v", err)
    }

    if *result != *expected {
        t.Errorf("Expected %+v, got %+v", expected, result)
    }
}
