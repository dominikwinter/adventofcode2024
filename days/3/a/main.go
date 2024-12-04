package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	context := ""

	for _, line := range input {
		context += line
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(context, -1)

	sum := 0

	for _, match := range matches {
		x := lib.ToInt(match[1])
		y := lib.ToInt(match[2])

		sum += x * y
	}

	return sum
}
