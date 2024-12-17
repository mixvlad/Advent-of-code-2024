package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Test input data
	input := `M.S
.A.
M.S`

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

	expected := 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestProcessFile2(t *testing.T) {
	// Test input data
	input := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

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

	expected := 9
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
