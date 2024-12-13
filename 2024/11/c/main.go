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

type CacheParams = [2]int
type CacheReturn = int
type CacheRecurse = ycache.Func[CacheParams, CacheReturn]

func Run(input string) any {
	res := 0
	calc := ycache.Y(ycache.Cache(calcFn))

	for _, c := range strings.Split(input, " ") {
		n := calc(CacheParams{lib.ToInt(c), 75})
		res += n
	}

	return res
}

func calcFn(recurse CacheRecurse) CacheRecurse {
	return func(params CacheParams) CacheReturn {
		num, n := params[0], params[1]

		if n == 0 {
			return 1
		}

		if num == 0 {
			return recurse(CacheParams{1, n - 1})
		}

		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			return 0 +
				recurse(CacheParams{lib.ToInt(str[:len(str)/2]), n - 1}) +
				recurse(CacheParams{lib.ToInt(str[len(str)/2:]), n - 1})
		}

		return recurse(CacheParams{num * 2024, n - 1})
	}
}
