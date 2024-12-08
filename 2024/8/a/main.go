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

type Point struct {
	x, y int
	c    string
}

func Run(input string) any {
	m := lib.StrToStrMatrix(input, "")

	fmt.Printf("%v\n\n", input)

	var points []Point

	for y, line := range m {
		for x, c := range line {
			if c != "." {
				fmt.Printf("%v %v %v\n", x, y, c)
				points = append(points, Point{x, y, c})
			}
		}
	}

	var antis []Point

	fmt.Printf("\n")

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 || p1.c != p2.c {
				continue
			}

			dx := p2.x - p1.x
			dy := p2.y - p1.y

			a1 := Point{x: p1.x - dx, y: p1.y - dy, c: "#"}
			a2 := Point{x: p2.x + dx, y: p2.y + dy, c: "#"}

			fmt.Printf("  a1: %v %v\n", a1, isInBounds(a1, m))
			if isInBounds(a1, m) {
				if !slices.Contains(antis, a1) {
					antis = append(antis, a1)
				}
			}

			fmt.Printf("  a2: %v %v\n", a2, isInBounds(a2, m))
			if isInBounds(a2, m) {
				if !slices.Contains(antis, a2) {
					antis = append(antis, a2)
				}
			}
		}
	}

	for y, line := range m {
		for x, c := range line {
			has := false
			for _, a := range antis {
				if a.x == x && a.y == y {
					fmt.Printf("%v", a.c)
					has = true
				}
			}
			if !has {
				fmt.Printf("%v", c)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n%v\n", antis)

	return len(antis)
}

func isInBounds(p Point, m [][]string) bool {
	return p.x >= 0 && p.x < len(m[0]) && p.y >= 0 && p.y < len(m)
}
