// Package common provides shared functionality for Advent of Code solutions
package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadNumberPairs reads pairs of numbers from an input file.
// Each line in the file should contain exactly two space-separated numbers.
// Returns two slices containing the first and second numbers from each pair.
func ReadNumberPairs(filepath string) ([]int, []int, error) {
	firstNums := make([]int, 0)
	secondNums := make([]int, 0)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file %s: %w", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		numbers := strings.Fields(line)

		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid input format at line %d: expected 2 numbers, got %d", lineNum, len(numbers))
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil {
			return nil, nil, fmt.Errorf("invalid first number at line %d: %s", lineNum, numbers[0])
		}
		if err2 != nil {
			return nil, nil, fmt.Errorf("invalid second number at line %d: %s", lineNum, numbers[1])
		}

		firstNums = append(firstNums, num1)
		secondNums = append(secondNums, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return firstNums, secondNums, nil
}

// AbsDiff returns the absolute difference between two integers
func AbsDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
