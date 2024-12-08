package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)
type machine struct {
	direction string
	row       int
	column    int
}

var directions = map[string][]int{
	"up":    {-1, 0},
	"right": {0, 1},
	"down":  {1, 0},
	"left": {0, -1},
}

func partOne(filename string) {
	grid, machine := readGrid(filename)
	exit := true
	sumVisited := 0
	for exit {
		direction := machine.direction
		row, column := machine.row, machine.column
		if row <= 0 || row >= len(grid) - 1 || column <= 0 || column >= len(grid[0]) - 1 {
			exit = false
			sumVisited += 1
			continue
		}
		// Check if we can move
		nextMove := grid[row+directions[direction][0]][column+directions[direction][1]]
		if nextMove == "." {
			grid[row+directions[direction][0]][column+directions[direction][1]] = "X"
			machine.row += directions[direction][0]
			machine.column += directions[direction][1]
			sumVisited += 1
		}
		if nextMove == "X" {
			machine.row += directions[direction][0]
			machine.column += directions[direction][1]
		}
		if nextMove == "#" {
			changeDirection(&machine, direction)
		}
	}

	log.Info("Sum of visited positions: ", sumVisited)
}

func changeDirection(machine *machine, direction string) {
	switch direction {
	case "up":
		machine.direction = "right"
	case "right":
		machine.direction = "down"
	case "down":
		machine.direction = "left"
	case "left":
		machine.direction = "up"
	}
}
func partTwo(filename string) {

}

func readGrid(filename string) ([][]string, machine) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create the grid
	fields := strings.Fields(string(content))
	grid := [][]string{}
	var m machine
	for idx, field := range fields {
		letters := strings.Split(field, "")
		rows := []string{}
		for idy, letter := range letters {
			if letter == "^" {
				m = machine{direction: "up", row: idx, column: idy}
			}
			rows = append(rows, letter)
		}
		grid = append(grid, rows)
	}
	return grid, m
}

func main() {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Part 1
	partOne("inputs/input_D6-E1.txt")
	partTwo("inputs/input_D6-E2.txt")
}