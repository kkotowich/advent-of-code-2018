package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Vector is a point on a gric
type Vector struct {
	x    int
	y    int
	size int
}

func main() {
	defer log.Printf("main took %s", time.Since(time.Now()))

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vectors := generateVectors(input)
	result1 := calcLargestArea(vectors)

	// closest, i := findClosestVector(vectors, 0, 0)

	// fmt.Println("---")
	// fmt.Println(string(closest.id))
	// fmt.Println(i)
	// fmt.Println("---")

	// for _, vector := range vectors {
	// 	fmt.Println(vector)
	// }
	// fmt.Println("-----")

	fmt.Println(result1)
}

func generateVectors(input []string) []Vector {
	var vectors []Vector
	for _, line := range input {
		x, y := parseLine(line)
		vectors = append(vectors, Vector{x, y, 0})
	}

	vectors = shiftVectors(vectors)

	return vectors
}

func shiftVectors(vectors []Vector) []Vector {
	xShift, yShift := findOriginVector(vectors)

	fmt.Println(xShift)
	fmt.Println(yShift)

	for i, vector := range vectors {
		vector.x -= xShift
		vector.y -= yShift
		vectors[i] = vector
	}

	return vectors
}

func findOriginVector(vectors []Vector) (int, int) {
	xShift := vectors[0].x
	yShift := vectors[0].y

	for _, vector := range vectors {
		if vector.x <= xShift {
			xShift = vector.x
		}
		if vector.y <= yShift {
			yShift = vector.y
		}
	}
	return xShift, yShift
}

func parseLine(input string) (int, int) {
	tokens := strings.Split(input, ", ")

	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])

	return x, y
}

func calcLargestArea(vectors []Vector) int {
	gridWidth, gridHeight := generateGrid(vectors)

	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			closestVector, closestIndex := findClosestVector(vectors, x, y)

			if closestIndex != -1 {
				closestVector.size++
				vectors[closestIndex] = closestVector
			}
		}
	}

	largestVector := vectors[0]
	for _, vector := range vectors {
		if !isInfinite(vector, gridWidth, gridHeight) {
			if vector.size > largestVector.size {
				largestVector = vector
			}
		}
	}

	return largestVector.size
}

func isInfinite(vector Vector, gridWidth, gridHeight int) bool {
	return vector.x == gridWidth || vector.y == gridHeight
}

func findClosestVector(vectors []Vector, targetX, targetY int) (Vector, int) {
	shortestDistance := 9999999
	var closestVector Vector
	var closestIndex int
	var collisions int

	for i, vector := range vectors {
		xDelta := 0
		yDelta := 0
		distance := 0

		if targetX > vector.x {
			xDelta = targetX - vector.x
		} else {
			xDelta = vector.x - targetX
		}

		if targetY > vector.y {
			yDelta = targetY - vector.y
		} else {
			yDelta = vector.y - targetY
		}
		distance = xDelta + yDelta
		if distance < shortestDistance {
			shortestDistance = distance
			closestVector = vector
			closestIndex = i
			collisions = 1
		} else if distance == shortestDistance {
			collisions++
		}
	}

	if collisions >= 2 {
		closestVector = Vector{0, 0, 0}
		closestIndex = -1
	}
	return closestVector, closestIndex
}

func generateGrid(vectors []Vector) (int, int) {
	x := 0
	y := 0

	for _, vector := range vectors {
		if vector.x > x {
			x = vector.x
		}
		if vector.y > y {
			y = vector.y
		}
	}

	return x, y
}
