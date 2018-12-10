package fileutil

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GetIntSliceFromFile gets []int from input file
func GetIntSliceFromFile(filename string) ([]int, error) {
	var result []int
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		token := scanner.Text()

		if token[0] == '+' {
			token = token[1:]
		}

		tokenInt, err := strconv.Atoi(token)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		result = append(result, tokenInt)
	}

	return result, nil
}

// GetStringSliceFromFile gets []string from input file
func GetStringSliceFromFile(filename string) ([]string, error) {

	var result []string
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

// GetInputFromFile gets the input from the file
func GetInputFromFile(filename string) []string {

	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
