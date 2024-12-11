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

	var calc BasicFunction
	calc = Y(func(f BasicFunction) BasicFunction {
		return func(num, n int) int {
			if n == 0 {
				return 1
			}

			if num == 0 {
				return calc(1, n-1)
			}

			str := strconv.Itoa(num)
			if len(str)%2 == 0 {
				return 0 +
					calc(lib.ToInt(str[:len(str)/2]), n-1) +
					calc(lib.ToInt(str[len(str)/2:]), n-1)
			}

			return calc(num*2024, n-1)
		}
	})

	for _, c := range strings.Split(input, " ") {
		n := calc(lib.ToInt(c), 75)
		res += n
	}

	return res
}

type BasicFunction func(int, int) int
type BasicFunctionTransformer func(BasicFunction) BasicFunction
type RecursiveFunctionToBasic func(RecursiveFunctionToBasic) BasicFunction

func Y(f BasicFunctionTransformer) BasicFunction {
	var cache = make(map[[2]int]int)

	g := func(h RecursiveFunctionToBasic) BasicFunction {
		return func(x, y int) int {
			key := [2]int{x, y}

			if ret, ok := cache[key]; ok {
				return ret
			}

			cache[key] = (f(h(h))(x, y))

			return cache[key]
		}
	}
	return g(g)
}
