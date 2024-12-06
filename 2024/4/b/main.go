package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	lines := lib.SplitLines(input)

	count := search(lines) // 0

	lines = lib.TransposeStrings(lines, 1)
	count += search(lines) // 90

	lines = lib.TransposeStrings(lines, 1)
	count += search(lines) // 180

	lines = lib.TransposeStrings(lines, 1)
	count += search(lines) // 270

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
