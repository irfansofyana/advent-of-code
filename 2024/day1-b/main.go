package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]int, []int, error) {
	firstArr := make([]int, 0)
	secondArr := make([]int, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return firstArr, secondArr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		if len(numbers) == 2 {
			num1, err1 := strconv.Atoi(numbers[0])
			num2, err2 := strconv.Atoi(numbers[1])

			if err1 == nil && err2 == nil {
				firstArr = append(firstArr, num1)
				secondArr = append(secondArr, num2)
			}

			if err1 != nil || err2 != nil {
				return firstArr, secondArr, errors.New("error parsing the numbers")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return firstArr, secondArr, err
	}

	return firstArr, secondArr, nil
}

func main() {
	firstArr, secondArr, err := readInput()
	if err != nil {
		log.Fatal("error reading input file")
	}

	mep := make(map[int]int)
	for _, num := range secondArr {
		mep[num]++
	}

	ans := 0
	for _, num := range firstArr {
		ans += num * mep[num]
	}

	log.Println("similarity score", ans)
}
