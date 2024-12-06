package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

type Map = [][]rune
type Pos struct {
	x int
	y int
	d rune
}

func Run(input string) any {
	var m Map = lib.StrToRuneMatrix(input, "")

	var p = getPos(&m)

	for {
		move(&m, p)
		// show(&m, p)

		if isOutOfMap(&m, p) {
			break
		}
	}

	return countX(&m) + 1
}

func isOutOfMap(m *Map, p *Pos) bool {
	return p.x == 0 || p.y == 0 || p.x == len((*m)[0])-1 || p.y == len(*m)-1
}

func getPos(m *Map) *Pos {
	for y, line := range *m {
		for x, d := range line {
			if d == '^' || d == 'v' || d == '<' || d == '>' {
				return &Pos{x, y, d}
			}
		}
	}

	return nil
}

func show(m *Map, p *Pos) {
	println("Pos:", p.x, p.y, string(p.d))
	for y, line := range *m {
		if y == 0 {
			fmt.Printf(" ")
			for x := range line {
				fmt.Printf("%v", x)
			}
			fmt.Printf("\n")
		}

		for x, c := range line {
			if x == 0 {
				fmt.Printf("%v", y)
			}
			fmt.Printf("%v", string(c))
		}
		fmt.Printf("\n")
	}
	println("")
}

func move(m *Map, p *Pos) {
	(*m)[p.y][p.x] = 'X'

	switch p.d {
	case '^':
		if (*m)[p.y-1][p.x] == '#' {
			p.d = '>'
		} else {
			p.y--
		}
	case 'v':
		if (*m)[p.y+1][p.x] == '#' {
			p.d = '<'
		} else {
			p.y++
		}
	case '<':
		if (*m)[p.y][p.x-1] == '#' {
			p.d = '^'
		} else {
			p.x--
		}
	case '>':
		if (*m)[p.y][p.x+1] == '#' {
			p.d = 'v'
		} else {
			p.x++
		}
	}

	(*m)[p.y][p.x] = p.d
}

func countX(m *Map) int {
	cnt := 0
	for _, line := range *m {
		for _, c := range line {
			if c == 'X' {
				cnt++
			}
		}
	}
	return cnt
}
