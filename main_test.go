package main

import (
	"bytes"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	bookPage := 1
	expectedLength := bookLength
	result := generateRandomString(bookPage)

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}

func TestSeekBook(t *testing.T) {
	result := seekBook([]string{"main.go", "10"})

	if len(result) != bookLength {
		t.Errorf("Expected length %d, but got %d", bookLength, len(result))
	}

	for _, char := range result {
		if !bytes.Contains([]byte(charSet), []byte{char}) {
			t.Errorf("Invalid character %c found in the generated string", char)
		}
	}
}

func TestSeekBookNoArg(t *testing.T) {
	result := seekBook([]string{"main.go"})

	if len(result) != 0 {
		t.Errorf("Expected length %d, but got %d", bookLength, len(result))
	}
}

func TestSeekBookNoInt(t *testing.T) {
	result := seekBook([]string{"main.go", "%"})

	if len(result) != 0 {
		t.Errorf("Expected length %d, but got %d", bookLength, len(result))
	}
}
