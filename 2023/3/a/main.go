package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

type Number struct {
	Value int
	Chars []rune
	Pos   []*lib.Point
}

type Symbol struct {
	Pos       lib.Point
	Neighbors []*Number
}

func Run(input string) any {
	isNum := false

	var numbers []*Number
	var symbols []*Symbol
	var num *Number

	for y, line := range lib.SplitLines(input) {
		for x, char := range line {
			if char >= '0' && char <= '9' {
				// start new number
				if !isNum {
					isNum = true

					num = &Number{
						Value: 0,
						Chars: []rune{},
						Pos:   []*lib.Point{},
					}
				}

				num.Pos = append(num.Pos, &lib.Point{X: x, Y: y})
				num.Chars = append(num.Chars, char)
			} else {
				// end number
				if isNum {
					isNum = false

					num.Value = lib.ToInt(string(num.Chars))
					numbers = append(numbers, num)
				}

				if char != '.' {
					symbols = append(symbols, &Symbol{
						Pos: lib.Point{X: x, Y: y},
					})
				}
			}
		}
	}

	println("numbers:")
	for _, n := range numbers {
		fmt.Printf("  %v\n", n)
	}

	println("gears:")
	for _, s := range symbols {
		fmt.Printf("  %v\n", s)
	}

	println("")

	sum := 0

	for _, n := range numbers {
		isNeighbor := false

		for _, s := range symbols {
			if isNeighbor {
				break
			}

			for _, p := range n.Pos {
				if s.Pos.IsNeighborDiagonal(p, 1) {
					fmt.Printf("  %v is neighbor of %v = %v %v\n", s.Pos, n.Value, p.X, p.Y)
					isNeighbor = true
					break
				}
			}
		}

		if isNeighbor {
			fmt.Printf("numbers with chars %v\n", n.Value)
			sum += n.Value
		}
	}

	return sum
}
