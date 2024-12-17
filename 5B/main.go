package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessFile handles the file reading and processing
func ProcessFile(filename string) (int, error) {
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

		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return -1, fmt.Errorf("Invalid line format: %s", line)
		}

		var key, value int
		fmt.Sscanf(parts[0], "%d", &key)
		fmt.Sscanf(parts[1], "%d", &value)

		rules[key] = append(rules[key], value)
	}

	// Function to check and fix sequence if needed
	processSequence := func(nums []int) ([]int, bool) {
		processed := make(map[int]bool)
		for i, num := range nums {
			if values, exists := rules[num]; exists {
				for _, value := range values {
					if processed[value] {
						// Found conflict, need to fix
						// Find position of the conflicting value
						var pos int
						for j := 0; j < i; j++ {
							if nums[j] == value {
								pos = j
								break
							}
						}
						// Create new sequence with current number moved before conflicting value
						newNums := make([]int, 0, len(nums))
						newNums = append(newNums, nums[:pos]...)
						newNums = append(newNums, num)
						newNums = append(newNums, nums[pos:i]...)
						newNums = append(newNums, nums[i+1:]...)
						return newNums, false
					}
				}
			}
			processed[num] = true
		}
		return nums, true
	}

	// Second part - processing numbers and fixing incorrect lines
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")

		var nums []int
		for _, numStr := range numbers {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			nums = append(nums, num)
		}

		wasFixed := false
		isValid := false
		attempts := 0
		maxAttempts := len(nums) * len(nums) // prevent infinite loops

		// Keep trying to fix sequence until it's valid or we hit max attempts
		for !isValid && attempts < maxAttempts {
			nums, isValid = processSequence(nums)
			if !isValid {
				wasFixed = true
			}
			attempts++
		}

		// Only count middle element if line needed fixing and was successfully fixed
		if wasFixed && isValid {
			middleIndex := len(nums) / 2
			total += nums[middleIndex]
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
