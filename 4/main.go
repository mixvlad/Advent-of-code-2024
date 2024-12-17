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
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check horizontal (right)
			if j <= cols-4 {
				word := string([]byte{grid[i][j], grid[i][j+1], grid[i][j+2], grid[i][j+3]})
				if word == "xmas" || word == "samx" {
					total++
				}
			}

			// Check vertical (down)
			if i <= rows-4 {
				word := string([]byte{grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j]})
				if word == "xmas" || word == "samx" {
					total++
				}

				// Check diagonal down-right
				if j <= cols-4 {
					word := string([]byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]})
					if word == "xmas" || word == "samx" {
						total++
					}
				}
			}

			// Check diagonal up-right
			if i >= 3 && j <= cols-4 {
				word := string([]byte{grid[i][j], grid[i-1][j+1], grid[i-2][j+2], grid[i-3][j+3]})
				if word == "xmas" || word == "samx" {
					total++
				}
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
