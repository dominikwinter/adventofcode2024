package days

import (
	"adventofcode2024/lib"
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func Day1A(scanner *bufio.Scanner) {
	lefts := make([]int, 0)
	rights := make([]int, 0)

	err := lib.Scan(scanner, func(line string) error {
		var err error
		var left int
		var right int

		cells := strings.Split(line, "   ")

		if left, err = strconv.Atoi(cells[0]); err != nil {
			return err
		}

		if right, err = strconv.Atoi(cells[1]); err != nil {
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

	slices.Sort(lefts)
	slices.Sort(rights)

	dist := 0

	for i := 0; i < len(lefts); i++ {
		dist += lib.Abs(lefts[i] - rights[i])
	}

	println(dist)
}
