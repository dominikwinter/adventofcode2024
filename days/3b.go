package days

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func Day3B(scanner *bufio.Scanner) {
	context := ""

	err := lib.Scan(scanner, func(line string) error {
		context += line
		return nil
	})

	if err != nil {
		println(err.Error())
	}

	re := regexp.MustCompile(`(do\(\)|don't\(\))|(mul\((\d{1,3}),(\d{1,3})\))`)
	matches := re.FindAllStringSubmatch(context, -1)

	do := true
	sum := 0

	for _, match := range matches {
		switch match[0] {
		case "do()":
			do = true
			continue
		case "don't()":
			do = false
			continue
		}

		if !do {
			continue
		}

		x, err := strconv.Atoi(match[3])
		if err != nil {
			println(err.Error())
		}

		y, err := strconv.Atoi(match[4])
		if err != nil {
			println(err.Error())
		}

		sum += x * y
	}

	fmt.Printf("%v\n", sum)
}
