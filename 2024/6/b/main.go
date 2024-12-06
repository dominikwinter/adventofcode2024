package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
	"slices"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

type Vector struct {
	X, Y int
}

type Point struct {
	X, Y int
	C    string
}

type Map [][]string

var dir = map[string]Vector{
	"^": {0, -1},
	"v": {0, 1},
	"<": {-1, 0},
	">": {1, 0},
}

var turn = map[string]string{
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

var debug = 0

func Run(input string) any {
	var err error

	m := lib.StrToStrMatrix(input, "")
	p := find(m)

	var infiniteObstacles []Vector

	for cnt := 0; cnt < 99999; cnt++ {
		p, err = next(p, m)

		if err != nil {
			break
		}

		println(cnt)

		// count only unique obstacles (x,y)
		if is, obstacle := infinity(p, m); is {
			if !slices.Contains(infiniteObstacles, obstacle) {
				infiniteObstacles = append(infiniteObstacles, obstacle)
			}
		}
	}

	return len(infiniteObstacles)
}

func infinity(p Point, mOrg Map) (bool, Vector) {
	m := duplicate(mOrg)
	obstacle, err := next(p, m)
	vec := Vector{obstacle.X, obstacle.Y}

	if err != nil {
		return false, vec
	}

	m[obstacle.Y][obstacle.X] = "O"

	for cnt := 0; cnt < 9999; cnt++ {
		p, err = next(p, m)

		if err != nil {
			return false, vec
		}
	}

	return true, vec // TODO check if we are in a loop, don't guess at limit
}

func duplicate(m Map) Map {
	m2 := make(Map, len(m))

	for y := range m {
		m2[y] = make([]string, len(m[y]))
		copy(m2[y], m[y])
	}

	return m2
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
	}

	// fmt.Printf("p=%v\n", p)

	p.X += d.X
	p.Y += d.Y

	return p, nil
}

func show(m Map) {
	fmt.Printf("\n===================================\n")

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			fmt.Printf("%v", m[y][x])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n===================================\n")
}

func find(m Map) Point {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := m[y][x]

			if c == "^" || c == ">" || c == "v" || c == "<" {
				return Point{x, y, c}
			}
		}
	}

	return Point{-1, -1, ""}
}
