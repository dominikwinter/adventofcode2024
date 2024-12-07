package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

type Equation struct {
	res  int
	nums []int
}

func Run(input string) any {
	lines := lib.SplitLines(input)
	eqs := make([]Equation, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ": ")

		num1, _ := strconv.Atoi(parts[0])
		num2 := lib.SliceToInts(strings.Split(parts[1], " "))

		eqs[i] = Equation{num1, num2}
	}

	sum := 0

	for _, e := range eqs {
		if calc(e.res, e.nums) {
			sum += e.res
		}
	}

	return sum
}

func calc(res int, num []int) bool {
	var _inner func(res int, num []int, i int, tmp int) bool

	_inner = func(res int, num []int, i int, tmp int) bool {
		if i >= len(num) {
			return tmp == res
		}

		next := num[i]

		if _inner(res, num, i+1, tmp+next) {
			return true
		}

		if _inner(res, num, i+1, tmp*next) {
			return true
		}

		sindog, _ := strconv.Atoi(strconv.Itoa(tmp) + strconv.Itoa(next))
		return _inner(res, num, i+1, sindog)
	}

	return _inner(res, num, 1, num[0])
}
