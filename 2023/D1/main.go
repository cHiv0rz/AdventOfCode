package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func findIndex(wordList []string, targetWord string) int {
	for i, word := range wordList {
		if word == targetWord {
			return i
		}
	}
	return -1 // Return -1 if the word is not found
}

func part_one(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := ""
		last := ""
		for _, char := range line {
			if unicode.IsDigit(char) && first == "" {
				first = string(char)
			} else if unicode.IsDigit(char) {
				last = string(char)
			}
		}
		if last == "" {
			code, _ := strconv.Atoi(first + first)
			fmt.Println(code)
			total += code
		} else {
			code, _ := strconv.Atoi(first + last)
			fmt.Println(code)
			total += code
		}
	}
	fmt.Printf("D1 - Part 1 - Final result: %d", total)
}
func part_two(scanner *bufio.Scanner) {
	wordList := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	total := 0
	// Create a regular expression pattern from the list of words
	pattern := fmt.Sprintf(`(?:%s)`, // \b for word boundaries
		regexp.QuoteMeta(wordList[0]),
	)

	for _, word := range wordList[1:] {
		pattern += fmt.Sprintf(`|(?:%s)`, regexp.QuoteMeta(word))
	}

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	for scanner.Scan() {
		line := scanner.Text()
		first := ""
		last := ""
		isWord := ""
		for index, char := range line {
			if unicode.IsDigit(char) {
				if index == 0 || first == "" {
					first = string(char)
				} else {
					last = string(char)
				}
				isWord = ""
			} else {
				isWord += string(char)
				if re.MatchString(isWord) {
					str := re.FindStringSubmatch(isWord)[0]
					index = findIndex(wordList, str)
					if first == "" {
						first = strconv.Itoa(index + 1)
					} else {
						last = strconv.Itoa(index + 1)
					}
					isWord = string(char)
				}
			}
		}
		fmt.Printf("For line %s, first is %s and last is %s\n", line, first, last)
		if last == "" {
			code, _ := strconv.Atoi(first + first)
			fmt.Println(code)
			total += code
		} else {
			code, _ := strconv.Atoi(first + last)
			fmt.Println(code)
			total += code
		}
	}
	fmt.Printf("Final result: %d", total)
}
func main() {
	// read a file from the command line
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	//part_one(scanner)
	part_two(scanner)
}
