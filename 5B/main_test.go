package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Test input data
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

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

	expected := 123
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
