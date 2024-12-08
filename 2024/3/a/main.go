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

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range matches {
		x := lib.ToInt(match[1])
		y := lib.ToInt(match[2])

		sum += x * y
	}

	return sum
}
