package main

import (
	"log"
	"sort"

	"common"
)

func main() {
	firstNums, secondNums, err := common.ReadNumberPairs("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	sort.Ints(firstNums)
	sort.Ints(secondNums)

	totalDiff := 0
	for i := 0; i < len(firstNums); i++ {
		totalDiff += common.AbsDiff(firstNums[i], secondNums[i])
	}

	log.Printf("Total absolute difference: %d", totalDiff)
}
