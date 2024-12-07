package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
	"slices"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

type Map [][]string

type Vector struct {
	X, Y int
}

type Point struct {
	X, Y int
	C    string
}

var dir = map[string]Vector{
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
}

var turn = map[string]string{
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

func Run(input string) any {
	m := lib.StrToStrMatrix(input, "")
	p, err := find(m)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	var cnt atomic.Uint32

	for y, line := range m {
		for x, c := range line {
			if c != "." {
				continue
			}

			wg.Add(1)

			go func(x, y int) {
				defer wg.Done()

				m2 := dup(m)
				m2[y][x] = "O"

				if isInfiniteLoop(p, m2) {
					cnt.Add(1)
				}
			}(x, y)
		}
	}

	wg.Wait()

	return int(cnt.Load())
}

func isInfiniteLoop(p Point, m Map) bool {
	var err error

	cache := make(map[Point]bool)

	// check by cache, but limit to 999999 steps, just to be on the safe side to
	// avoid an infinite loop
	for cnt := 0; cnt < 999999; cnt++ {
		p, err = next(p, m)

		if err != nil {
			return false
		}

		if cache[p] {
			return true
		}

		cache[p] = true
	}

	return true
}

func dup(m Map) Map {
	d := make(Map, len(m))

	for y, line := range m {
		d[y] = slices.Clone(line)
	}

	return d
}

func next(p Point, m Map) (Point, error) {
	d := dir[p.C]
	x := p.X + d.X
	y := p.Y + d.Y

	if y < 0 || y >= len(m) || x < 0 || x >= len(m[0]) {
		return Point{}, fmt.Errorf("out of bounds")
	}

	if m[y][x] == "#" || m[y][x] == "O" {
		p.C = turn[p.C]
		d = dir[p.C]
	} else {
		p.X += d.X
		p.Y += d.Y
	}

	return p, nil
}

func find(m Map) (Point, error) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := m[y][x]

			if c == "^" || c == ">" || c == "v" || c == "<" {
				return Point{x, y, c}, nil
			}
		}
	}

	return Point{}, fmt.Errorf("not found")
}
