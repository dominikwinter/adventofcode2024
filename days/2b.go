package days

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Day2B(scanner *bufio.Scanner) {
	safeCnt := 0

	err := lib.Scan(scanner, func(line string) error {
		fmt.Printf("%v\n", line)

		cells := make([]int, 0)

		// split by space, convert to int and append to cells
		for _, cell := range strings.Split(line, " ") {
			num, err := strconv.Atoi(cell)
			if err != nil {
				return err
			}
			cells = append(cells, num)
		}

		// default check
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

		return nil
	})

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Printf("%v\n", safeCnt)
}

func check(cells []int) bool {
	fmt.Printf("  %v\n", cells)

	var dir string

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
