package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)
type Rule struct {
	Left int
	Right int
}
type Update []int

func partOne(filename string) {
	rules, updates, err := parseInput(filename)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}
	result := sumCorrectlyOrderedUpdates(rules, updates)
	log.Info("Sum of all middle page numbers correctly ordered: ", result)
}
func partTwo(filename string) {
	rules, updates, err := parseInput(filename)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}
	result := sumFixedUpdates(rules, updates)
	log.Info("Sum of all middle page numbers fixed: ", result)
}

func isValidOrder(update Update, rules []Rule) bool {
	// Create a map of positions
	positions := make(map[int]int)
	for i, page := range update {
		positions[page] = i
	}
	// Check each rule
	for _, rule := range rules {
		// Skip rules that don't apply to this update
		beforePos, beforeExists := positions[rule.Left]
		afterPos, afterExists := positions[rule.Right]
	
	// If both pages exist in the update, check their order
		if beforeExists && afterExists {
			if beforePos > afterPos {
				return false
			}
		}
	}
	return true
}

func parseInput(filename string) ([]Rule, []Update, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules []Rule
	var updates []Update
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule{Left: before, Right: after})
		} else if strings.Contains(line, ",") {
			var update Update
			parts := strings.Split(line, ",")
			for _, part := range parts {
				number, _ := strconv.Atoi(part)
				update = append(update, number)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

func sumCorrectlyOrderedUpdates(rules []Rule, updates []Update) int {
	sum := 0
	for _, update := range updates {
		if isValidOrder(update, rules) {
			middleIndex := len(update) / 2
			sum += update[middleIndex]
		}
	}
	return sum
}
func sumFixedUpdates(rules []Rule, updates []Update) int {
	sum := 0
	for _, update := range updates {
		if !isValidOrder(update, rules) {
			// Make a copy of the update that we can modify
			fixed := make(Update, len(update))
			copy(fixed, update)
			// Keep swapping until the order is valid
			changed := true
			for changed {
				changed = false
				// Try to fix order by swapping adjacent elements
				for i := 0; i < len(fixed)-1; i++ {
					if needsSwap(fixed[i], fixed[i+1], rules) {
						fixed[i], fixed[i+1] = fixed[i+1], fixed[i]
						changed = true
					}
				}
			}
			
			// Get middle number from fixed update
			middleIndex := len(fixed) / 2
			sum += fixed[middleIndex]
		}
	}
	return sum
}

func needsSwap(after int, before int, rules []Rule) bool {
    for _, rule := range rules {
        if rule.Left == before && rule.Right == after {
            return true
        }
    }
    return false
}

func main() {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Part 1
	partOne("inputs/input_D5-E1.txt")
	partTwo("inputs/input_D5-E2.txt")
}