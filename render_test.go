package main

import (
	"os"
	"strings"
	"testing"
)

func loadTestBanner(t *testing.T) []string {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Fatalf("failed to read banner: %v", err)
	}
	return strings.Split(string(data), "\n")
}

func TestEmptyInput(t *testing.T) {
	banner := loadTestBanner(t)
	result := Render("", banner, "", "")

	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestSingleCharacter(t *testing.T) {
	banner := loadTestBanner(t)
	result := Render("A", banner, "", "")

	lines := strings.Split(result, "\n")

	if len(lines) < 8 {
		t.Errorf("expected at least 8 lines, got %d", len(lines))
	}
}

func TestNewLineInput(t *testing.T) {
	banner := loadTestBanner(t)
	result := Render("\\n", banner, "", "")

	if result != "\n" {
		t.Errorf("expected newline, got %q", result)
	}
}

func TestHello(t *testing.T) {
	banner := loadTestBanner(t)
	result := Render("Hello", banner, "", "")

	if result == "" {
		t.Errorf("expected non-empty output")
	}
}

func TestHelloWithNewLine(t *testing.T) {
	banner := loadTestBanner(t)
	result := Render("Hello\\nThere", banner, "", "")

	parts := strings.Split(result, "\n\n")

	if len(parts) < 2 {
		t.Errorf("expected two blocks separated by newline")
	}
}

func TestMultipleNewLines(t *testing.T) {
	banner := loadTestBanner(t)
	result := Render("Hello\\n\\nThere", banner, "", "")

	if !strings.Contains(result, "\n\n") {
		t.Errorf("expected empty line between blocks")
	}
}
