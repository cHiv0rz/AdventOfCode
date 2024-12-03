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
		re := regexp.MustCompile(`\d+,\d+`)
		parts := re.FindString(match)
		numbers := strings.Split(parts, ",")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		sum += num1 * num2
	}
	log.Info("Sum of all multiplications: ", sum)
}
func partTwo(filename string) {

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
	//partTwo("inputs/input_D1-E2.txt")
}