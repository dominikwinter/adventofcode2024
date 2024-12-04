package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"os"
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
	count := search(input)

	input = lib.TransposeStrings(input, 1)
	count += search(input)

	input = lib.TransposeStrings(input, 1)
	count += search(input)

	input = lib.TransposeStrings(input, 1)
	count += search(input)

	return count
}

func search(input []string) int {
	count := 0

	for y := 0; y < len(input)-2; y++ {
		for x := 0; x < len(input[y])-2; x++ {
			if string(input[y][x]) == "M" && string(input[y][x+2]) == "S" &&
				string(input[y+1][x+1]) == "A" &&
				string(input[y+2][x]) == "M" && string(input[y+2][x+2]) == "S" {
				count++
			}
		}
	}

	return count
}
