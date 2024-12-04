package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	result := Run(input)

	fmt.Printf("%v\n", result)
}

func Run(input []string) any {
	const WORD = "XMAS"

	count := search(input, WORD)

	inputD1 := lib.TransposeDiagonalStrings(input)
	count += search(inputD1, WORD)

	input = lib.TransposeStrings(input, 1)
	count += search(input, WORD)

	inputD2 := lib.TransposeDiagonalStrings(input)
	count += search(inputD2, WORD)

	return count
}

func search(input []string, word string) int {
	count := 0

	for _, line := range input {
		count += strings.Count(line, word) + strings.Count(line, lib.Reverse(word))
	}

	return count
}
