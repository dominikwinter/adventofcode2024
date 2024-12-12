package main

import (
	"fmt"
	"os"
	"time"
	"unsafe"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

type Plant struct {
	*lib.Node
	letter string
	group  int64
}

type Map map[int]map[int]*Plant
type Groups map[int64]map[*Plant]bool

func Run(input string) any {
	matrix := lib.ToMatrix(input, "", func(y, x int, letter string) *Plant {
		node := &lib.Node{Y: y, X: x}

		return &Plant{
			Node:   node,
			letter: letter,
			group:  *(*int64)(unsafe.Pointer(&node)),
		}
	})

	groups := make(Groups)

	for _, line := range matrix {
		for _, current := range line {
			fillGroups(current, matrix, groups)
		}
	}

	for node, nodes := range groups {
		fmt.Printf("id:%v %v\n", node, len(nodes))
	}

	return 0
}

func fillGroups(start *Plant, m Map, groups Groups) {
	if _, ok := groups[start.group]; !ok {
		groups[start.group] = make(map[*Plant]bool)
	}

	for _, d := range lib.Dir {
		y := start.Y + d[0]
		x := start.X + d[1]

		if neighbor, ok := m[y][x]; ok {
			if _, ok := groups[start.group][neighbor]; ok {
				continue
			}

			if neighbor.letter == start.letter {
				neighbor.group = start.group
				groups[start.group][neighbor] = true
				fillGroups(neighbor, m, groups)
			}
		}
	}
}
