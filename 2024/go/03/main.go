package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type match struct {
	pos  int
	text string
}

func findSolution(input string) {
	// for the first solution
	// combinedRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)

	// for the second solution
	combinedRegex := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	var matches []match
	allMatches := combinedRegex.FindAllStringSubmatch(input, -1)
	allMatchesIndex := combinedRegex.FindAllStringSubmatchIndex(input, -1)

	for i := range allMatches {
		matches = append(matches, match{
			pos:  allMatchesIndex[i][0],
			text: allMatches[i][0],
		})
	}

	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].pos < matches[j].pos
	})

	for _, m := range matches {
		fmt.Println(m.pos, m.text)
	}

	shouldMultiply := true
	var sum int64 = 0
	for _, m := range matches {
		if m.text == "do()" {
			shouldMultiply = true
		} else if m.text == "don't()" {
			shouldMultiply = false
		} else {
			numMatches := regexp.MustCompile(`\d+`).FindAllString(m.text, -1)
			if len(numMatches) != 2 {
				fmt.Printf("Warning: Invalid mul pattern found: %s\n", m.text)
				continue
			}

			num1, err := strconv.Atoi(numMatches[0])
			if err != nil {
				fmt.Printf("Error parsing first number in %s: %v\n", m.text, err)
				continue
			}

			num2, err := strconv.Atoi(numMatches[1])
			if err != nil {
				fmt.Printf("Error parsing second number in %s: %v\n", m.text, err)
				continue
			}

			if shouldMultiply {
				product := int64(num1) * int64(num2)
				sum += product
			}
		}
	}

	fmt.Println("Total sum is", sum)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	findSolution(string(content))
}
