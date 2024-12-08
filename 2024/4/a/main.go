package main

import (
	"fmt"
	"os"
	"strings"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	const WORD = "XMAS"

	lines := lib.SplitLines(input)

	count := search(lines, WORD)

	inputD1 := lib.TransposeDiagonalStrings(lines)
	count += search(inputD1, WORD)

	lines = lib.TransposeStrings(lines, 1)
	count += search(lines, WORD)

	inputD2 := lib.TransposeDiagonalStrings(lines)
	count += search(inputD2, WORD)

	return count
}

func search(input []string, word string) int {
	drow := lib.Reverse(word)
	count := 0

	for _, line := range input {
		count += strings.Count(line, word) + strings.Count(line, drow)
	}

	return count
}
