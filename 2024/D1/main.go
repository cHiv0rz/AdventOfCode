package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func partOne(filename string) {

	//Open file
	data, err := os.Open(filename)
	if err != nil {
		log.WithError(err).Error("ERROR opening file")
		return
	}
	defer data.Close()

	//Define variables
	scanner := bufio.NewScanner(data)
	var firstColumn []int
	var secondColumn []int
	var differences int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Info("Skipping invalid line:", line)
			continue
		}
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Error("Error converting first column to integer:", err)
			continue
		}
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Error("Error converting first column to integer:", err)
			continue
		}

		// Append to the slices
		firstColumn = append(firstColumn, first)
		secondColumn = append(secondColumn, second)
	}
	sort.Ints(firstColumn)
	sort.Ints(secondColumn)

	for idx := 0; idx < len(firstColumn); idx++ {
		distance := secondColumn[idx] - firstColumn[idx]
		if distance < 0 {
			distance *= -1
		}
		differences += distance
	}
	log.Info("The difference in distances is: ", differences)
}
func partTwo(filename string) {
	//Open file
	data, err := os.Open(filename)
	if err != nil {
		log.WithError(err).Error("ERROR opening file")
		return
	}
	defer data.Close()

	//Define variables
	scanner := bufio.NewScanner(data)
	var firstColumn []int
	var secondColumn []int

	var similarityScore int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Info("Skipping invalid line:", line)
			continue
		}
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Error("Error converting first column to integer:", err)
			continue
		}
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Error("Error converting first column to integer:", err)
			continue
		}

		// Append to the slices
		firstColumn = append(firstColumn, first)
		secondColumn = append(secondColumn, second)
	}

	// TODO: Improve this to not iterate n*n times
	for idx1 := 0; idx1 < len(firstColumn); idx1++ {
		times := 0
		for idx2 := 0; idx2 < len(secondColumn); idx2++ {
			if firstColumn[idx1] == secondColumn[idx2] {
				times++
			}
		}
		similarityScore += firstColumn[idx1] * times
	}
	log.Info("The similarity score is: ", similarityScore)
}
func main() {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Part 1
	partOne("inputs/input_D1-E1.txt")
	partTwo("inputs/input_D1-E2.txt")
}