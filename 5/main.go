package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessFile handles the file reading and processing
func ProcessFile(filename string) (int, error) {
	// Read file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1, err
	}
	defer file.Close()

	// Create hash map and total variable
	rules := make(map[int][]int)
	total := 0

	scanner := bufio.NewScanner(file)
	// First part - reading rules
	for scanner.Scan() {
		line := scanner.Text()

		// Break if empty line
		if line == "" {
			break
		}

		// Split line by '|'
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return -1, fmt.Errorf("Invalid line format: %s", line)
		}

		// Parse key and value
		var key, value int
		fmt.Sscanf(parts[0], "%d", &key)
		fmt.Sscanf(parts[1], "%d", &value)

		// Add value to existing array or create new one
		rules[key] = append(rules[key], value)
	}

	// Second part - processing numbers
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")

		// Create empty set for processed numbers
		processed := make(map[int]bool)

		// Flag to track if line is correct
		lineIsCorrect := true

		// Convert string numbers to integers
		var nums []int
		for _, numStr := range numbers {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			nums = append(nums, num)
		}

		// Process each number in line
		for _, num := range nums {
			// Find values in hashtable
			if values, exists := rules[num]; exists {
				// Check if any of the values are in processed set
				for _, value := range values {
					if processed[value] {
						lineIsCorrect = false
						break
					}
				}
			}
			if !lineIsCorrect {
				break
			}
			// Add current number to processed set
			processed[num] = true
		}

		// If line is correct, add middle element to total
		if lineIsCorrect {
			middleIndex := len(nums) / 2
			middleNum := nums[middleIndex]
			total += middleNum
		}
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return total, nil
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
