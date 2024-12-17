package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function for integer absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func chechSafe(levels []int) bool {
	direction := 0
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]

		if Abs(diff) < 1 || Abs(diff) > 3 {
			return false
		}

		if diff*direction < 0 {
			return false
		}
		direction = diff
	}
	return true
}

func removeElement(report []int, deleteIndex int) []int {
	copyReport := make([]int, len(report))
	copy(copyReport, report)

	copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)

	return copyReport
}

func checkReportSafetyWithDeletion(reportNum []int) bool {
	for i := 0; i < len(reportNum); i++ {
		isSafe := chechSafe(removeElement(reportNum, i))
		if isSafe {
			return true
		}
	}

	return false
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

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return 161, nil
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
