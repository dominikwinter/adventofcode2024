package days

import (
	"adventofcode2024/lib"
	"bufio"
	"regexp"
	"strconv"
)

func Day3A(scanner *bufio.Scanner) {
	context := ""

	err := lib.Scan(scanner, func(line string) error {
		context += line
		return nil
	})

	if err != nil {
		println(err.Error())
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(context, -1)

	sum := 0

	for _, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			println(err.Error())
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			println(err.Error())
		}

		sum += x * y
	}

	println(sum)
}
