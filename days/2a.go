package days

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Day2A(scanner *bufio.Scanner) {
	safeCnt := 0

	err := lib.Scan(scanner, func(line string) error {
		fmt.Printf("line: %v\n", line)

		cells := strings.Split(line, " ")
		row := make([]int, len(cells))

		fmt.Printf("cells: %v\n", cells)

		var dir string

		isSafe := true

		for i, cell := range cells {
			num, err := strconv.Atoi(cell)

			if err != nil {
				return err
			}

			row[i] = num

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

		return nil
	})

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Printf("%v\n", safeCnt)
}
