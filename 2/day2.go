package main

import (
	fileutil "code-advent/utils"
	"fmt"
)

func main() {

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	counts := getCounts(input)
	result1 := calcCheckSum(counts)
	fmt.Println(result1)

	result2 := calcAnswer2(input)
	fmt.Println(result2)
}

func getCounts(input []string) []map[rune]int {
	var counts []map[rune]int

	for i := 0; i < len(input); i++ {
		tokenCounts := make(map[rune]int)
		for _, char := range input[i] {
			tokenCounts[char]++
		}
		counts = append(counts, tokenCounts)
	}
	return counts
}

func calcCheckSum(input []map[rune]int) int {
	if len(input) == 0 {
		return 0
	}

	count := []int{0, 0}

	for _, wordMap := range input {

		wordCount := []int{0, 0}

		for _, value := range wordMap {
			if value == 2 {
				wordCount[0] = 1
			} else if value == 3 {
				wordCount[1] = 1
			}
		}
		count[0] += wordCount[0]
		count[1] += wordCount[1]
	}

	return count[0] * count[1]
}

func calcAnswer2(input []string) string {
	var result string
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			hDistance, indices := calcHammingDistance(input[i], input[j])
			if hDistance == 1 {
				return input[i][:indices[0]] + input[i][indices[0]+1:]
			}
		}
	}
	return result
}

func calcHammingDistance(value1, value2 string) (int, []int) {
	hDistance := 0
	indices := []int{}
	for key := range value1 {
		if value1[key] != value2[key] {
			hDistance++
			indices = append(indices, key)
		}
	}
	return hDistance, indices
}
