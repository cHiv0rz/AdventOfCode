package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)
func partOne(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var safeReports int

	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)
		safe := true
		ascending := false
		descending := false
		if len(levels) > 10 {
			log.Info("Skipping invalid line:", line)
			continue
		}
		first, err := strconv.Atoi(levels[0])
		if err != nil {
			log.Error("Error converting level to integer:", err)
			continue
		}
		second, err := strconv.Atoi(levels[1])
		if err != nil {
			log.Error("Error converting level to integer:", err)
			continue
		}
		if first == second {
			safe = false
			continue
		}
		if first < second {
			if second - first > 3 {
				continue
			}
			ascending = true
		} else if second < first {
			if first - second > 3 {
				continue
			}
			descending = true
		}
		for idx := 1; idx < len(levels); idx++{
			if idx + 1 == len(levels) {
				continue
			}
			value, err := strconv.Atoi(levels[idx])
			if err != nil {
				log.Error("Error converting level to integer:", err)
				continue
			}
			nextValue, err := strconv.Atoi(levels[idx + 1])
			if err != nil {
				log.Error("Error converting level to integer:", err)
				continue
			}
			if value == nextValue {
				safe = false
				break
			}
			if ascending {
				if nextValue < value {
					safe = false
					break
				}
				if nextValue - value > 3 {
					safe = false
					break
				}
			} else if descending {
				if nextValue > value {
					safe = false
					break
				}
				if value - nextValue > 3 {
					safe = false
					break
				}
			}
		}
		if safe {
			safeReports++
		}
	}
	log.Info("D2-E1: Number of safe reports: ", safeReports)
}

func partTwo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var safeReports int
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		numbers := fieldsToInts(fields)
		if len(numbers) > 10 {
			log.Info("Skipping invalid line:", line)
			continue
		}
		if IsSafeReportWithoutOne(numbers) {
			safeReports++
		}
	}
	log.Info("D2-E2: Number of safe reports: ", safeReports)
}

func IsSafeUpReport(report []int) bool {
	isSafe := true
	for idx := 0; idx < len(report)-1; idx++ {
		if 1 > (report[idx+1]-report[idx]) || (report[idx+1]-report[idx]) > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func IsSafeDownReport(report []int) bool {
	isSafe := true
	for idx := 0; idx < len(report)-1; idx++ {
		if 1 > (report[idx]-report[idx+1]) || (report[idx]-report[idx+1]) > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func IsSafeReport(report []int) bool {
	return IsSafeUpReport(report) || IsSafeDownReport(report)
}

func IsSafeReportWithoutOne(report []int) bool {
	isSafe := false
	for idx := 0; idx < len(report); idx++ {
		if IsSafeReport(RemoveElementAt(idx, report)) {
			isSafe = true
			break
		}
	}
	return isSafe
}

func RemoveElementAt[T any](index int, report []T) []T {
	if index < 0 || index >= len(report) {
		return []T{}
	}

	temp := make([]T, 0)
	for idx, value := range report {
		if idx != index {
			temp = append(temp, value)
		}
	}

	return temp
}

func fieldsToInts(fields []string) []int {
    numbers := make([]int, 0, len(fields))
    for _, f := range fields {
        num, err := strconv.Atoi(f)
        if err != nil {
            continue
        }
        numbers = append(numbers, num)
    }
    return numbers
}

func main() {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Part 1
	partOne("inputs/input_D2-E1.txt")
	// Part 2
	partTwo("inputs/input_D2-E2.txt")
}
