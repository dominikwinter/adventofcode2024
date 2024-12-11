package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func Run(input string) any {
	res := 0
	cc := cacheCalc()

	for _, c := range strings.Split(input, " ") {
		n := cc(lib.ToInt(c), 75)
		res += n
	}

	return res
}

func cacheCalc() func(int, int) int {
	var cache = make(map[[2]int]int)
	var calc func(int, int) int

	calc = func(num int, n int) int {
		if n == 0 {
			return 1
		}

		key := [2]int{num, n}

		if ret, ok := cache[key]; ok {
			return ret
		}

		var result int

		if num == 0 {
			result = calc(1, n-1)
		} else {
			str := strconv.Itoa(num)
			if len(str)%2 == 0 {
				result = 0 +
					calc(lib.ToInt(str[:len(str)/2]), n-1) +
					calc(lib.ToInt(str[len(str)/2:]), n-1)
			} else {
				result = calc(num*2024, n-1)
			}
		}

		cache[key] = result

		return result
	}

	return calc
}
