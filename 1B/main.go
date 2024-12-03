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

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create slices to store numbers
	var firstColumn []int
	secondColumn := make(map[int]int)

	// Create scanner to read file
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// Split line by whitespace
		numbers := strings.Fields(line)

		// Convert strings to integers and append to appropriate slices
		if len(numbers) == 2 {
			num1, err1 := strconv.Atoi(numbers[0])
			num2, err2 := strconv.Atoi(numbers[1])

			if err1 == nil && err2 == nil {
				firstColumn = append(firstColumn, num1)
				secondColumn[num2]++
			}
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0

	for i := 0; i < len(firstColumn); i++ {
		result += firstColumn[i] * secondColumn[firstColumn[i]]
	}

	// Print results
	fmt.Println("Result:", result)
}
