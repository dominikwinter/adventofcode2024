package main

import (
	"fmt"
	"os"
	"strings"
	"time"

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
	group  *lib.Node
}

type Group struct {
	letter string
	nodes  map[*Plant]bool
	fences int
}

type Matrix map[int]map[int]*Plant
type Groups map[*lib.Node]*Group

func Run(input string) any {
	matrix := lib.ToMatrix(input, "", func(y, x int, letter string) *Plant {
		node := &lib.Node{Y: y, X: x}
		return &Plant{
			Node:   node,
			letter: letter,
			group:  node,
		}
	})

	// show(matrix)

	groups := make(Groups)

	for _, line := range matrix {
		for _, current := range line {
			fillGroups(current, matrix, groups)
		}
	}

	for _, group := range groups {
		for current := range group.nodes {
			countFences(current, matrix, group)
		}
	}

	sum := 0
	for node, group := range groups {
		fmt.Printf("l:%v id:%v nodes:%v fences:%v mul:%v\n", group.letter, node, len(group.nodes), group.fences, len(group.nodes)*group.fences)
		sum += len(group.nodes) * group.fences
	}

	return sum
}

func show(matrix map[int]map[int]*Plant) {
	fmt.Print("    |")
	for x := 0; x < len(matrix[0]); x++ {
		fmt.Printf("%3v", x)
	}
	fmt.Println("\n----+" + strings.Repeat("-", 3*len(matrix[0])))

	for y := 0; y < len(matrix); y++ {
		fmt.Printf("%3v |", y)
		for x := 0; x < len(matrix[y]); x++ {
			fmt.Printf("%3v", matrix[y][x].letter)
		}
		fmt.Println()
	}

	fmt.Println()
}

func fillGroups(plant *Plant, matrix Matrix, groups Groups) {
	// fmt.Printf("node:%v (y:%v,x:%v)\n", node.letter, node.Y, node.X)

	// init group
	if _, ok := groups[plant.group]; !ok {
		groups[plant.group] = &Group{
			letter: plant.letter,
			nodes:  map[*Plant]bool{plant: true},
			fences: 0,
		}
	}

	// check neighbors
	for _, d := range lib.Dir {
		y := plant.Y + d[0]
		x := plant.X + d[1]

		if neighbor, ok := matrix[y][x]; ok {
			if neighbor.letter == plant.letter {
				if _, ok := groups[plant.group].nodes[neighbor]; !ok {
					// fmt.Printf("  neighbor:%v %v -> add to group\n", neighbor.letter, neighbor.Node)
					neighbor.group = plant.group
					groups[plant.group].nodes[neighbor] = true
					fillGroups(neighbor, matrix, groups)
				}
			}
		}
	}
}

func countFences(plant *Plant, matrix Matrix, group *Group) {
	// check neighbors
	for _, d := range lib.Dir {
		y := plant.Y + d[0]
		x := plant.X + d[1]

		if neighbor, ok := matrix[y][x]; ok {
			if neighbor.letter != plant.letter {
				group.fences++
				// fmt.Printf("node:%v %v %v %v\n", plant.letter, plant.Node, neighbor.Node, group.fences)
			}
		} else {
			group.fences++
			// fmt.Printf("node:%v %v %v\n", plant.letter, plant.Node, group.fences)
		}
	}
}
