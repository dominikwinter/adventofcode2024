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

		x := lib.ToInt(match[3])
		y := lib.ToInt(match[4])

		sum += x * y
	}

	return sum
}
