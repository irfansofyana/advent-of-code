package main

import (
	"bufio"
	"common"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindSolution(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %s: %w", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	cntSafe := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		numbers := strings.Fields(line)

		if isSafe(numbers) {
			cntSafe++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return cntSafe, nil
}

func FindSolution2Bruteforce(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %s: %w", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	cntSafe := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		numbers := strings.Fields(line)

		if isSafe(numbers) {
			cntSafe++
			continue
		}

		for i := 0; i < len(numbers); i++ {
			tmp := make([]string, 0, len(numbers)-1)
			tmp = append(tmp, numbers[:i]...)
			tmp = append(tmp, numbers[i+1:]...)
			if isSafe(tmp) {
				cntSafe++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return cntSafe, nil
}

func isSafe(numbers []string) bool {
	prv := 0
	isIncreasing := false
	isSafe := true

	for i := 0; i < len(numbers) && isSafe; i++ {
		num, _ := strconv.Atoi(numbers[i])

		if i == 0 {
			prv = num
			continue
		}

		if prv == num {
			isSafe = false
			break
		}

		if prv < num && i == 1 {
			isIncreasing = true
			if !isAcceptableDiff(prv, num) {
				isSafe = false
			}
			prv = num
			continue
		}
		if prv > num && i == 1 {
			if !isAcceptableDiff(prv, num) {
				isSafe = false
			}
			prv = num
			continue
		}

		if prv < num && !isIncreasing {
			isSafe = false
			break
		}

		if prv > num && isIncreasing {
			isSafe = false
			break
		}

		if !isAcceptableDiff(prv, num) {
			isSafe = false
			break
		}

		prv = num
	}

	return isSafe
}

func isAcceptableDiff(a, b int) bool {
	return common.AbsDiff(a, b) >= 1 && common.AbsDiff(a, b) <= 3
}

func main() {
	cntDiff, err := FindSolution2Bruteforce("input.txt")
	if err != nil {
		log.Fatalf("Failed to find solution: %v", err)
	}

	log.Printf("safe reports: %d", cntDiff)
}
