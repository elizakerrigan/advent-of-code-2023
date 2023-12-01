package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type CalibrationLine struct {
	Text string
}

func (cl CalibrationLine) Value() int {
	firstDigit, lastDigit := -1, -1

	for _, r := range cl.Text {
		if unicode.IsDigit(r) {
			if firstDigit == -1 {
				firstDigit = int(r - '0')
			}
			lastDigit = int(r - '0')
		}
	}

	if firstDigit != -1 && lastDigit != -1 {
		value, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		return value
	}

	return 0
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

	for scanner.Scan() {
		line := scanner.Text()
		calibrationLine := CalibrationLine{Text: line}
		sum += calibrationLine.Value()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading from file:", err)
	}

	fmt.Println("Total sum:", sum)
}
