package main

import (
	"strings"
	"testing"
)

func TestQuote(t *testing.T) {
	quote := prepareQuotes("smart_quotes.txt")
	if len(quote) <= 10 {
		t.Errorf("Quotes was too short, got: %d, want at least %d.", len(quote), 10)
	}
}

func TestFirstQuote(t *testing.T) {
	quote := prepareQuotes("smart_quotes.txt")
	firstQuote := quote[0]

	if !strings.Contains(firstQuote, "achieve") {
		t.Errorf("First quote should contain the word \"achieve\".")
	}
}
