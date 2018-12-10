package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Vector starting point of a claim
type Vector struct {
	x int
	y int
}

// Claim a claim of area
type Claim struct {
	id       int
	position Vector
	width    int
	height   int
}

func main() {

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	start := time.Now()

	result1, result2 := calcTotalOverlapCount(input)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	fmt.Println(result1)
	fmt.Println(result2)
}

func calcTotalOverlapCount(input []string) (int, int) {
	var claim Claim
	var claims []Claim
	claimCounts := make(map[Vector]int)
	collisionCount := 0
	var heroClaimID int

	for _, value := range input {
		claim = parseClaim(value)
		claims = append(claims, claim)

		for x := claim.position.x; x < claim.position.x+claim.width; x++ {
			for y := claim.position.y; y < claim.position.y+claim.height; y++ {
				claimCounts[Vector{x, y}]++
				if claimCounts[Vector{x, y}] == 2 {
					collisionCount++
				}
			}
		}
	}

	for _, claim := range claims {
		isHero := true
		for x := claim.position.x; x < claim.position.x+claim.width; x++ {
			for y := claim.position.y; y < claim.position.y+claim.height; y++ {
				if claimCounts[Vector{x, y}] != 1 {
					isHero = false
				}
			}
		}

		if isHero {
			heroClaimID = claim.id
			break
		}
	}

	return collisionCount, heroClaimID
}

func claimArea(claimedArea [][]int, claim Claim) [][]int {
	for i := claim.position.x; i < (claim.position.x + claim.width); i++ {
		for j := claim.position.y; j < (claim.position.y + claim.height); j++ {
			claimedArea[i][j]++
		}
	}
	return claimedArea
}

func parseClaim(input string) Claim {
	// #1 @ 1,3: 4x4
	tokens := strings.Split(input, " ")

	id := parseID(tokens[0])
	position := parsePosition(tokens[2])
	width, height := parseSize(tokens[3])

	return Claim{id, position, width, height}
}

func parseID(input string) int {
	id, _ := strconv.Atoi(input[1:])
	return id
}

func parsePosition(input string) Vector {
	positionArray := strings.Split(input[:len(input)-1], ",")

	x, _ := strconv.Atoi(positionArray[0])
	y, _ := strconv.Atoi(positionArray[1])

	return Vector{x, y}
}

func parseSize(input string) (int, int) {
	sizeArray := strings.Split(input, "x")

	w, _ := strconv.Atoi(sizeArray[0])
	h, _ := strconv.Atoi(sizeArray[1])

	return w, h
}
