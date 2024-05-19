package logging

import (
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestNewSFDailyLogger(t *testing.T) {
	tempFile, err := os.CreateTemp("", "log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	logger, closer, err := NewSFDailyLogger(tempFile.Name(), slog.LevelInfo)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer closer.Close()

	if logger == nil {
		t.Fatal("Expected logger to be not nil")
	}

	testLog := "This is a test log"
	logger.Info(testLog)

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	println(string(content))

	if !strings.Contains(string(content), testLog) {
		t.Fatalf("Expected log file to contain: %s", testLog)
	}
}
