package main

import (
	"fmt"
	"os"
	"slices"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	m := lib.StrToIntMatrix(input, "   ")
	m = lib.Transpose(m, 1)

	slices.Sort(m[0])
	slices.Sort(m[1])

	res := 0

	for i, n := range m[0] {
		res += lib.Abs(n - m[1][i])
	}

	return res
}
