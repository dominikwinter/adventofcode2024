package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	lib "github.com/dominikwinter/adventofcode2024/lib"
	"github.com/dominikwinter/adventofcode2024/lib/ycache"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func Run(input string) any {
	res := 0

	calc := ycache.Y(ycache.Cache(calcFn))

	for _, c := range strings.Split(input, " ") {
		n := calc([2]int{lib.ToInt(c), 75})
		res += n
	}

	return res
}

type CFnP = [2]int

func calcFn(recurse ycache.Func[CFnP, int]) ycache.Func[CFnP, int] {
	return func(params CFnP) int {
		num, n := params[0], params[1]

		if n == 0 {
			return 1
		}

		if num == 0 {
			return recurse(CFnP{1, n - 1})
		}

		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			return 0 +
				recurse(CFnP{lib.ToInt(str[:len(str)/2]), n - 1}) +
				recurse(CFnP{lib.ToInt(str[len(str)/2:]), n - 1})
		}

		return recurse(CFnP{num * 2024, n - 1})
	}
}
