package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Test input data
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

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
	safeLevels, err := ProcessFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	expected := 4
	if safeLevels != expected {
		t.Errorf("Expected %d safe levels, but got %d", expected, safeLevels)
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{-1, 1},
		{0, 0},
		{-5, 5},
		{10, 10},
	}

	for _, test := range tests {
		result := Abs(test.input)
		if result != test.expected {
			t.Errorf("Abs(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
