package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"time"
	"unicode"
)

func main() {
	defer log.Printf("main took %s", time.Since(time.Now()))

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result1 := startChainReaction(input[0])
	result2 := calcUnitForBestEffeciency(input[0])

	fmt.Println(len(result1))
	fmt.Println(result2)
}

func startChainReaction(input string) string {
	var done bool
	result := []rune(input)

	for !done {
		done = true
		for i := len(result) - 1; i > 0; i-- {
			if unicode.ToUpper(result[i]) == unicode.ToUpper(result[i-1]) &&
				result[i] != result[i-1] {

				if i >= len(result)-1 {
					result = append(result[0:i-1], result[i:]...)
				} else {
					result = append(result[0:i-1], result[i+1:]...)
				}
				done = false
			}
		}
	}

	return string(result)
}

func calcUnitForBestEffeciency(input string) int {
	units := "abcdefghijklmnopqrstuvwqyz"
	shortestLength := len(input)

	for _, unit := range units {
		experimentalPolymer := removeAllUnits(input, unit)
		result := len(startChainReaction(experimentalPolymer))

		if result < shortestLength {
			shortestLength = result
		}
	}

	return shortestLength
}

func removeAllUnits(input string, unit rune) string {
	result := []rune(input)

	for i := len(result) - 1; i > 0; i-- {
		if unicode.ToUpper(result[i]) == unicode.ToUpper(unit) {
			result = append(result[0:i], result[i+1:]...)
		}
	}

	return string(result)
}
