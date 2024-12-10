package main

import (
	"fmt"
	"os"
	"time"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func Run(input string) any {
	m := lib.StrToIntMatrix(input, "")

	paths := &[][]lib.Node{}

	show(m)

	for y, line := range m {
		for x, step := range line {
			if step == 0 {
				path := &[]lib.Node{}

				climb(&lib.Node{Y: y, X: x}, step+1, &m, path)

				*paths = append(*paths, *path)
			}
		}
	}

	score := 0
	for _, nodes := range *paths {
		fmt.Printf("%v\n", nodes)
		score += len(nodes)
	}

	return score
}

func climb(node *lib.Node, step int, m *[][]int, paths *[]lib.Node) {
	for _, neighbor := range node.GetNeighbors(m) {
		if neighbor.GetValue(m) == step {
			if step == 9 {
				// lib.UniqueAdd(paths, *neighbor)
				*paths = append(*paths, *neighbor)
			} else {
				climb(neighbor, step+1, m, paths)
			}
		}
	}
}

func show(m [][]int) {
	for _, line := range m {
		for _, n := range line {
			fmt.Printf("%3d", n)
		}
		fmt.Println()
	}
}
