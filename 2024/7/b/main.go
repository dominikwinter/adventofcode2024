package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func Run(input string) any {
	sum := 0

	for _, line := range lib.SplitLines(input) {
		parts := strings.Split(line, ": ")

		res := lib.ToInt(parts[0])
		nums := lib.SliceToInts(strings.Split(parts[1], " "))

		if calc(res, nums) {
			sum += res
		}
	}

	return sum
}

func calc(res int, nums []int) bool {
	stack := []int{nums[0]}

	for _, next := range nums[1:] {
		var newStack []int

		for _, val := range stack {
			if add := val + next; add != res || add > res {
				newStack = append(newStack, add)
			} else {
				return true
			}

			if mul := val * next; mul != res || mul > res {
				newStack = append(newStack, mul)
			} else {
				return true
			}

			if con := lib.ToInt(strconv.Itoa(val) + strconv.Itoa(next)); con != res {
				newStack = append(newStack, con)
			} else {
				return true
			}
		}

		stack = newStack
	}

	return false
}
