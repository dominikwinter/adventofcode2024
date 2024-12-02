package days

import (
	"adventofcode2024/lib"
	"bufio"
	"strconv"
	"strings"
)

func Day1B(scanner *bufio.Scanner) {
	lefts := make([]int, 0)
	rights := make([]int, 0)

	err := lib.Scan(scanner, func(line string) error {
		cells := strings.Split(line, "   ")

		left, err := strconv.Atoi(cells[0])
		if err != nil {
			return err
		}

		right, err := strconv.Atoi(cells[1])
		if err != nil {
			return err
		}

		lefts = append(lefts, left)
		rights = append(rights, right)

		return nil
	})

	if err != nil {
		println(err.Error())
		return
	}

	result := 0

	for _, l := range lefts {
		for _, r := range rights {
			if l == r {
				result += l
			}
		}
	}

	println(result)
}
