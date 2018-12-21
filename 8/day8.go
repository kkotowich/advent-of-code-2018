package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Node is a node on the tree
type Node struct {
	header     Header
	childNodes []Node
	metadata   []int
}

func (n Node) sumMetadata() int {
	var sum int
	for _, value := range n.metadata {
		sum += value
	}
	for _, childNode := range n.childNodes {
		sum += childNode.sumMetadata()
	}
	return sum
}

func (n Node) sumMetadata2() int {
	var sum int

	if len(n.childNodes) > 0 {
		for _, value := range n.metadata {
			if !(value <= 0 || value > len(n.childNodes)) {
				sum += n.childNodes[value-1].sumMetadata2()
			}
		}
	} else {
		for _, value := range n.metadata {
			sum += value
		}
	}

	return sum
}

// Header contains counts for child nodes and metadata
type Header struct {
	childNodeCount int
	metadataCount  int
}

func main() {
	defer log.Printf("main took %s", time.Since(time.Now()))

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tokens := strings.Split(input[0], " ")
	rootNode, _ := buildTree(tokens, 0)
	result1 := rootNode.sumMetadata()
	result2 := rootNode.sumMetadata2()

	fmt.Println(result1)
	fmt.Println(result2)
}

func buildTree(tokens []string, index int) (Node, int) {
	var childNode Node
	var tailIndex int
	var childNodes []Node

	// fmt.Println(index)

	childCount, _ := strconv.Atoi(tokens[index])
	metadataCount, _ := strconv.Atoi(tokens[index+1])
	nextIndex := index + 2

	// fmt.Println(childCount)

	for i := 1; i <= childCount; i++ {
		childNode, tailIndex = buildTree(tokens, nextIndex)

		nextIndex = tailIndex

		childNodes = append(childNodes, childNode)
	}

	var metadata []int
	for i := 0; i < metadataCount; i++ {
		value, _ := strconv.Atoi(tokens[nextIndex+i])
		metadata = append(metadata, value)

		// fmt.Println(value)
	}

	// fmt.Println(metadataCount)

	nextIndex += metadataCount
	node := Node{Header{childCount, metadataCount}, childNodes, metadata}

	return node, nextIndex
}
