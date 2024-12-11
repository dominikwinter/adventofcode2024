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
	cache := cacheCalc()

	for _, c := range strings.Split(input, " ") {
		n := cache(lib.ToInt(c), 75)
		res += n
	}

	return res
}

type CacheCalcFn = func(int, int) int

func cacheCalc() CacheCalcFn {
	var cache = make(map[[2]int]int)
	var wrapper CacheCalcFn
	var calc CacheCalcFn

	wrapper = func(num int, n int) int {
		key := [2]int{num, n}

		if ret, ok := cache[key]; ok {
			return ret
		}

		cache[key] = calc(num, n)
		return cache[key]
	}

	calc = func(num, n int) int {
		if n == 0 {
			return 1
		}

		if num == 0 {
			return wrapper(1, n-1)
		}

		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			return 0 +
				wrapper(lib.ToInt(str[:len(str)/2]), n-1) +
				wrapper(lib.ToInt(str[len(str)/2:]), n-1)
		}

		return wrapper(num*2024, n-1)
	}

	return calc
}
