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
	return _inner(res, nums, 1, nums[0])
}

func _inner(res int, nums []int, i int, tmp int) bool {
	if i >= len(nums) {
		return tmp == res
	}

	next := nums[i]

	if _inner(res, nums, i+1, tmp+next) {
		return true
	}

	if _inner(res, nums, i+1, tmp*next) {
		return true
	}

	return _inner(res, nums, i+1, lib.ToInt(strconv.Itoa(tmp)+strconv.Itoa(next)))
}
