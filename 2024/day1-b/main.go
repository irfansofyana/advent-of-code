package main

import (
	"log"

	"common"
)

// calculateSimilarityScore computes the similarity score between two arrays
// by counting the frequency of numbers in the second array and multiplying
// each number in the first array by its frequency in the second array
func calculateSimilarityScore(first, second []int) int {
	// Create frequency map for the second array
	frequencies := make(map[int]int)
	for _, num := range second {
		frequencies[num]++
	}

	// Calculate similarity score
	score := 0
	for _, num := range first {
		score += num * frequencies[num]
	}
	return score
}

func main() {
	firstNums, secondNums, err := common.ReadNumberPairs("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	score := calculateSimilarityScore(firstNums, secondNums)
	log.Printf("Similarity score: %d", score)
}
