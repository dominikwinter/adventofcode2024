package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"os"
	"strings"
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
	l1 := make([]int, len(input))
	l2 := make([]int, len(input))

	for i, line := range input {
		n := lib.SliceToInts(strings.Split(line, "   "))

		l1[i] = n[0]
		l2[i] = n[1]
	}

	res := 0

	for _, l := range l1 {
		for _, r := range l2 {
			if l == r {
				res += l
			}
		}
	}

	return res
}
