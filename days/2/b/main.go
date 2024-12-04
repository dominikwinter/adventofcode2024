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
	safeCnt := 0

	for _, line := range input {
		fmt.Printf("%v\n", line)

		cells := lib.SliceToInts(strings.Split(line, " "))
		save := check(cells)

		// tolerate check: Problem Dampener
		if !save {
			// check by removing one by one cell
			for i := range cells {
				save = check(lib.Remove(cells, i))

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

func check(cells []int) bool {
	fmt.Printf("  %v\n", cells)

	var dir string

	// loop to second last element
	for i := 0; i < len(cells)-1; i++ {
		curr := cells[i]
		next := cells[i+1]
		diff := next - curr

		fmt.Printf("	Debug: %v | %v - %v = %v\n", i, next, curr, diff)

		// check keeps in order
		if i == 0 {
			if diff > 0 {
				dir = "inc"
			} else {
				dir = "dec"
			}
		} else {
			if (dir == "inc" && diff < 0) || (dir == "dec" && diff > 0) {
				fmt.Printf("	Invalid sequence: changed direction: %v\n", dir)
				return false
			}
		}

		abs := lib.Abs(diff)

		if abs == 0 {
			fmt.Printf("	Invalid sequence: neither an increase or a decrease: %v\n", abs)
			return false
		}

		if abs > 3 {
			fmt.Printf("	Invalid sequence: diff to big: %v\n", abs)
			return false
		}
	}

	return true
}
