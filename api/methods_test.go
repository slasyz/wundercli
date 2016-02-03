package api

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	expectedTaskText := "Do something"
	expectedDate := fmt.Sprintf("%d-05-12", time.Now().Year())

	taskText, date := parseDate("Do something     May 12")
	if taskText != expectedTaskText {
		t.Fatalf("Wrong task text: expected \"%s\", got \"%s\"", expectedTaskText, taskText)
	}
	if date != expectedDate {
		t.Fatalf("Wrong month: expected %s, got %s", expectedDate, date)
	}
}
