package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)
func partOne(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Error("Error reading file: ", err)
		return
	}
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(string(content), -1)
	sum := 0
	for _, match := range matches {
		numbers := convertToNumbers(match)
		sum += numbers[0] * numbers[1]
	}
	log.Info("D3,Part1 - Sum of all multiplications: ", sum)
}
func partTwo(filename string) {
	content, err := os.ReadFile(filename)
	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	if err != nil {
		log.Error("Error reading file: ", err)
		return
	}
	sum := 0
	enabled := true
	matches := re.FindAllString(string(content), -1)
	for _, match := range matches {
		if strings.Contains(match, "mul") {
			if enabled {
				numbers := convertToNumbers(match)
				sum += numbers[0] * numbers[1]
			} else {
				continue
			}
		} else if strings.Contains(match, "don't()") {
			enabled = false
		} else if strings.Contains(match, "do()") {
			enabled = true
		}
	}
	
	log.Info("D3,Part2 - Sum of all multiplications: ", sum)
}
func convertToNumbers(match string) []int {
	re := regexp.MustCompile(`\d+,\d+`)
	parts := re.FindString(match)
	numbers := strings.Split(parts, ",")
	num1, _ := strconv.Atoi(numbers[0])
	num2, _ := strconv.Atoi(numbers[1])
	return []int{num1, num2}
}
func main() {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Part 1
	partOne("inputs/input_D3-E1.txt")
	// Part 2
	partTwo("inputs/input_D3-E2.txt")
}