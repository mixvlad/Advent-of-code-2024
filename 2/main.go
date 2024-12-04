package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function for integer absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ProcessFile handles the file reading and processing
func ProcessFile(filename string) (int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line
	safeLevels := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into numbers
		numbersStr := strings.Fields(line)

		// Convert strings to integers
		levels := make([]int, 0, len(numbersStr))
		for _, numStr := range numbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			levels = append(levels, num)
		}

		safeLevel := true
		direction := 0
		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i] - levels[i+1]

			if Abs(diff) < 1 || Abs(diff) > 3 {
				safeLevel = false
				break
			}

			if diff*direction < 0 {
				safeLevel = false
				break
			}
			direction = diff
		}

		if safeLevel {
			safeLevels++
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeLevels, nil
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	safeLevels, err := ProcessFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Safe levels: %d", safeLevels)
}
