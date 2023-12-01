package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numberMap = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func convertToDigit(s string) int {
	if val, ok := numberMap[s]; ok {
		return val
	}
	digit, _ := strconv.Atoi(s)
	return digit
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	digitRegex := regexp.MustCompile(`(?:\d|one|two|three|four|five|six|seven|eight|nine)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := digitRegex.FindAllString(line, -1)
		if len(matches) >= 2 {
			first, last := matches[0], matches[len(matches)-1]
			firstDigit, lastDigit := convertToDigit(first), convertToDigit(last)
			value, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
			sum += value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading from file:", err)
	}

	fmt.Println("Total sum:", sum)
}
