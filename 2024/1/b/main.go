package main

import (
	"fmt"
	"os"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	m := lib.StrToIntMatrix(input, "   ")
	m = lib.Transpose(m, 1)

	res := 0

	for _, l := range m[0] {
		for _, r := range m[1] {
			if l == r {
				res += l
			}
		}
	}

	return res
}
