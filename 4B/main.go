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

	// Read lines into array
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return -1, err
	}

	total := 0
	rows := len(grid)
	if rows == 0 {
		fmt.Println("Empty file")
		return -1, err
	}
	cols := len(grid[0])

	// Single loop to check all directions
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {

			word1 := string([]byte{grid[i-1][j-1], grid[i][j], grid[i+1][j+1]})
			word2 := string([]byte{grid[i+1][j-1], grid[i][j], grid[i-1][j+1]})
			if (word1 == "mas" && word2 == "mas") || (word1 == "sam" && word2 == "sam") || (word1 == "mas" && word2 == "sam") || (word1 == "sam" && word2 == "mas") {
				total++
			}
		}
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
