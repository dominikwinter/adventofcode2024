package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	input = strings.ReplaceAll(input, "\n", "")

	re := regexp.MustCompile(`(do\(\)|don't\(\))|(mul\((\d{1,3}),(\d{1,3})\))`)
	matches := re.FindAllStringSubmatch(input, -1)

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

		x := lib.ToInt(match[3])
		y := lib.ToInt(match[4])

		sum += x * y
	}

	return sum
}
