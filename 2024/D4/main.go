package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)
func partOne(filename string) {
	grid := readGrid(filename)
	sum := 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == "X" {
				// Check forwards
				if column + 3 <= len(grid[0]) - 1 {
					if grid[row][column+1] == "M" && grid[row][column+2] == "A" && grid[row][column+3] == "S" {
						sum++
					}
				}
				// Check backwards
				if column - 3 >= 0 {
					if grid[row][column-1] == "M" && grid[row][column-2] == "A" && grid[row][column-3] == "S" {
						sum++
					}
				}
				// Check downwards
				if row + 3 <= len(grid) - 1 {
					if grid[row+1][column] == "M" && grid[row+2][column] == "A" && grid[row+3][column] == "S" {
						sum++
					}
				}
				// Check upwards
				if row - 3 >= 0 {
					if grid[row-1][column] == "M" && grid[row-2][column] == "A" && grid[row-3][column] == "S" {
						sum++
					}
				}
				// Check if can go diagional top-left
				if row - 3 >= 0 && column - 3 >= 0 {
					if grid[row-1][column-1] == "M" && grid[row-2][column-2] == "A" && grid[row-3][column-3] == "S" {
						sum++
					} 
				}
				// Check if can go diagional bottom-left
				if row + 3 <= len(grid) - 1 && column - 3 >= 0 {
					if grid[row+1][column-1] == "M" && grid[row+2][column-2] == "A" && grid[row+3][column-3] == "S" {
						sum++
					} 
				}
				// Check if can go diagional top-right
				if row - 3 >= 0 && column + 3 <= len(grid[0]) - 1 {
					if grid[row-1][column+1] == "M" && grid[row-2][column+2] == "A" && grid[row-3][column+3] == "S" {
						sum++
					} 
				}
				// Check if can go diagional bottom-right
				if row + 3 <= len(grid) - 1 && column + 3 <= len(grid[0]) - 1 {
					if grid[row+1][column+1] == "M" && grid[row+2][column+2] == "A" && grid[row+3][column+3] == "S" {
						sum++
					} 
				}
			}				
		}
	}
	log.Info("D4,Part1 - Sum of all XMAS occurrences: ", sum)

}
func partTwo(filename string) {
	grid := readGrid(filename)
	sum := 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == "A" {
				exists, err := checkIfMAS(&grid, row, column)
				if err != nil {
					continue
				}
				if exists {
					sum++
				}
			}
		}
	}
	log.Info("D4,Part2 - Sum of all X-MAS occurrences: ", sum)
}

func checkIfMAS(grid *[][]string, row int, column int) (bool, error) {
	exists := false
	if row - 1 < 0 || row + 1 > len(*grid) - 1 {
		return exists, nil
	}
	if column - 1 < 0 || column + 1 > len((*grid)[0]) - 1 {
		return exists, nil
	}
	
	joined1 := (*grid)[row-1][column-1] + (*grid)[row+1][column+1]
	joined2 := (*grid)[row+1][column-1] + (*grid)[row-1][column+1]

	log.Info("joined1: ", joined1, " - joined2: ", joined2)

	if (joined1 == "MS" && joined2 == "MS") || (joined1 == "SM" && joined2 == "SM") ||
		(joined1 == "MS" && joined2 == "SM") || (joined1 == "SM" && joined2 == "MS") {
		exists = true
	}
	return exists, nil
}

func readGrid(filename string) [][]string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create the grid
	fields := strings.Fields(string(content))
	grid := [][]string{}
	for _, field := range fields {
		letters := strings.Split(field, "")
		rows := []string{}
		for _, letter := range letters {
			rows = append(rows, letter)
		}
		grid = append(grid, rows)
	}
	return grid
}
func main() {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Part 1
	partOne("inputs/input_D4-E1.txt")
	partTwo("inputs/input_D4-E2.txt")
}
