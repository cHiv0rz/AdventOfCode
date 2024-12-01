package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)
func partOne(filename string) {

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
	partOne("inputs/input_D1-E1.txt")
	partTwo("inputs/input_D1-E2.txt")
}