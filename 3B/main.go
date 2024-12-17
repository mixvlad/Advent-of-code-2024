package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processText(text string) int {
	// Regular expression to find mul(X,Y) patterns
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	// Find all matches
	matches := re.FindAllString(text, -1)

	// Process each match
	result := 0
	for _, match := range matches {
		matchResult, err := multiply(match)
		if err == nil {
			result += matchResult
		}
	}

	return result
}

func multiply(input string) (int, error) {
	// Regular expression to extract numbers
	re := regexp.MustCompile(`\d{1,3}`)

	// Find matches
	matches := re.FindAllString(input, -1)
	if matches == nil {
		return 0, fmt.Errorf("invalid input format")
	}

	// Convert strings to integers
	x, err := strconv.Atoi(matches[0])
	if err != nil {
		return 0, err
	}

	y, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	return x * y, nil
}

// ProcessFile handles the file reading and processing
func ProcessFile(filename string) (int, error) {
	// Read the entire file content into a byte slice
	data, err := os.ReadFile(filename)
	if err != nil {
		return -1, err
	}

	initText := string(data)

	// Regex pattern (escaped properly for Go)
	pattern := `don\'t\(\)(?s).*?(do\(\)|$)`

	re := regexp.MustCompile(pattern)

	nonMatches := re.Split(initText, -1)

	result := strings.Join(nonMatches, "\n")

	return processText(result), nil
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	result, err := ProcessFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Print(result)
}
