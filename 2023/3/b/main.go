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

type Gear struct {
	Pos       lib.Point
	Neighbors []*Number
}

func Run(input string) any {
	isNum := false

	var numbers []*Number
	var gears []*Gear
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

				if char == '*' {
					gears = append(gears, &Gear{
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
	for _, g := range gears {
		fmt.Printf("  %v\n", g)
	}

	println("")

	for _, g := range gears {
		for _, n := range numbers {
			isNeighbor := false

			for _, p := range n.Pos {
				if g.Pos.IsNeighborDiagonal(p, 1) {
					isNeighbor = true
					break
				}
			}

			if isNeighbor {
				g.Neighbors = append(g.Neighbors, n)

				fmt.Printf("neighbor of %v\n", n.Value)
			}
		}
	}

	println("")

	sum := 0

	for _, g := range gears {
		if len(g.Neighbors) < 2 {
			continue
		}

		mul := 1
		for _, n := range g.Neighbors {
			mul *= n.Value
		}

		sum += mul
	}

	return sum
}
