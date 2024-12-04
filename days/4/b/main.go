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
	count := search(input) // 0

	input = lib.TransposeStrings(input, 1)
	count += search(input) // 90

	input = lib.TransposeStrings(input, 1)
	count += search(input) // 180

	input = lib.TransposeStrings(input, 1)
	count += search(input) // 270

	return count
}

/**
 * Search in matrix this pattern:
 *
 *   M * S
 *   * A *
 *   M * S
 *
 * where * could be any character
 */
func search(input []string) int {
	count := 0

	for y := 0; y < len(input)-2; y++ {
		for x := 0; x < len(input[y])-2; x++ {
			if input[y][x] == 'M' && input[y][x+2] == 'S' &&
				input[y+1][x+1] == 'A' &&
				input[y+2][x] == 'M' && input[y+2][x+2] == 'S' {
				count++
			}
		}
	}

	return count
}
