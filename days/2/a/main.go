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
		var dir string
		fmt.Printf("line: %v\n", line)

		cells := strings.Split(line, " ")
		row := make([]int, len(cells))
		isSafe := true

		fmt.Printf("cells: %v\n", cells)

		for i, cell := range cells {
			row[i] = lib.ToInt(cell)

			if i > 0 {
				diff := row[i] - row[i-1]

				if i == 1 {
					if diff > 0 {
						dir = "inc"
					} else {
						dir = "dec"
					}
				}

				if i > 1 {
					if (dir == "inc" && diff < 0) || (dir == "dec" && diff > 0) {
						isSafe = false
						fmt.Printf("Invalid sequence: changed direction: %v - %v\n", diff, dir)
						break
					}
				}

				abs := lib.Abs(diff)

				if abs == 0 {
					isSafe = false
					fmt.Printf("Invalid sequence: neither an increase or a decrease: %v\n", abs)
					break
				}

				if abs > 3 {
					isSafe = false
					fmt.Printf("Invalid sequence: diff to big: %v\n", abs)
					break
				}
			}
		}

		if isSafe {
			safeCnt++
		}

		fmt.Printf("safe: %v - %v\n", isSafe, safeCnt)
	}

	return safeCnt
}
