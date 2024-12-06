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
	failed := 0
	lines := lib.StrToIntMatrix(input, " ")

	for _, line := range lines {
		fmt.Printf("line: %v\n", line)

		asc := lib.GetSorted(line)
		desc := lib.GetReverse(asc)

		for i, curr := range line {
			if curr != asc[i] && curr != desc[i] {
				failed++
				fmt.Printf("invalid sequence: %v\n", curr)
				break
			}

			if i > 0 {
				diff := line[i] - line[i-1]

				if diff == 0 {
					failed++
					fmt.Printf("invalid sequence: neither an increase or a decrease\n")
					break
				}

				abs := lib.Abs(diff)

				if abs > 3 {
					failed++
					fmt.Printf("invalid sequence: diff to big: %v\n", abs)
					break
				}
			}
		}

		fmt.Printf("failed: %v\nsafe: %v\n\n", failed, len(lines)-failed)
	}

	return len(lines) - failed
}
