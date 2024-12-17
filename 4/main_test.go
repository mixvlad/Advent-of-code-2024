package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Test input data
	input := `..X...
.SAMX.
.A..A.
XMAS.S
.X....`

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

	expected := 4
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestProcessFile2(t *testing.T) {
	// Test input data
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

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

	expected := 18
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
