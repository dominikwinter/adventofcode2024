package main

import (
	"fmt"
	"os"
	"slices"
	"time"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

type Vector struct {
	lib.Node
}

type Antenna struct {
	lib.Node
	freq string
}

type Antinode struct {
	lib.Node
}

func Run(input string) any {
	m := lib.StrToStrMatrix(input, "")

	var nodes []Antenna

	for y, line := range m {
		for x, c := range line {
			if c != "." {
				fmt.Printf("%v %v %v\n", x, y, c)
				nodes = append(nodes, Antenna{lib.Node{X: x, Y: y}, c})
			}
		}
	}

	var antis []Antinode

	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1 == n2 || n1.freq != n2.freq {
				continue
			}

			d := Vector{lib.Node{X: n2.X - n1.X, Y: n2.Y - n1.Y}}

			tmp1 := Antinode{n1.Node}
			for {
				if !tmp1.Node.Backward(&d.Node).IsInBounds(&m) {
					break
				}

				if !slices.Contains(antis, tmp1) {
					antis = append(antis, tmp1)
				}
			}

			tmp2 := Antinode{n2.Node}
			for {
				// happy accident, i don't know. actually a copy paste error.
				// in fact i wanted to use Forward() here
				if !tmp2.Node.Backward(&d.Node).IsInBounds(&m) {
					break
				}

				if !slices.Contains(antis, tmp2) {
					antis = append(antis, tmp2)
				}
			}
		}
	}

	return len(antis)
}
