package main

import (
	"fmt"
	"os"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	safeCnt := 0
	lines := lib.StrToIntMatrix(input, " ")

	for _, line := range lines {
		fmt.Printf("%v\n", line)

		save := check(line)

		// tolerate check: Problem Dampener
		if !save {
			// check by removing one by one cell
			for i := range line {
				save = check(lib.Remove(line, i))

				if save {
					break
				}
			}
		}

		if save {
			safeCnt++
		}
	}

	return safeCnt
}

func check(line []int) bool {
	fmt.Printf("line: %v\n", line)

	asc := lib.GetSorted(line)
	desc := lib.GetReverse(asc)

	for i, curr := range line {
		if curr != asc[i] && curr != desc[i] {
			fmt.Printf("invalid sequence: %v\n", curr)
			return false
		}

		if i > 0 {
			diff := line[i] - line[i-1]

			if diff == 0 {
				fmt.Printf("invalid sequence: neither an increase or a decrease\n")
				return false
			}

			abs := lib.Abs(diff)

			if abs > 3 {
				fmt.Printf("invalid sequence: diff to big: %v\n", abs)
				return false
			}
		}
	}

	return true
}
