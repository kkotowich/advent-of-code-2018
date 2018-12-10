package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()

	input, err := fileutil.GetIntSliceFromFile("input.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result1 := calcSum(input)
	fmt.Println(result1)
	fmt.Println("")

	result2 := findDuplicateSumPartial(input)
	fmt.Println("")
	fmt.Println(result2)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func calcSum(input []int) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

func findDuplicateSumPartial(input []int) int {
	var foundSums []int
	sum := 0

	whileCount := 0
	forCount := 0

	for true {
		whileCount++
		for i := 0; i < len(input); i++ {
			forCount++
			foundSums = append(foundSums, sum)
			sum += input[i]
			if checkSum(foundSums, sum) {
				fmt.Println(whileCount)
				fmt.Println(forCount)
				return sum
			}
		}
	}

	return 0
}

func checkSum(haystack []int, needle int) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}
