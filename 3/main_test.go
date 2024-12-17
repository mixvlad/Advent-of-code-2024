package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Test input data
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	// Create a temporary test file
	tmpfile, err := os.CreateTemp("", "test_input.*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test data to the file
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test the ProcessFile function
	result, err := ProcessFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	expected := 161
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
