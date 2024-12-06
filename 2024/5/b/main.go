package main

import (
	"adventofcode2024/lib"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
}

func Run(input string) any {
	sections := strings.Split(input, "\n\n")
	before, after := parseRules(sections[0])
	orders := lib.StrToIntMatrix(sections[1], ",")

	var corrected []int

	for pos, order := range orders {
		fmt.Printf("\n=> %v\n", order)

		l := len(order)
		for i := 0; i < l; i++ {
			curr := order[i]
			b := before[curr]
			a := after[curr]

			if i < l-1 {
				next := order[i+1]
				fmt.Printf("  is next %v in %v\n", next, a)

				if slices.Contains(a, next) {
					fmt.Printf("  = ERROR, yes %v\n", next)
					corrected = append(corrected, pos)
					orders[pos][i], orders[pos][i+1] = orders[pos][i+1], orders[pos][i] // swap with next
					i = -1
				}
			}

			if i > 0 {
				prev := order[i-1]
				fmt.Printf("  is prev %v in %v\n", prev, b)

				if slices.Contains(b, prev) {
					fmt.Printf("  = ERROR, yes %v\n", prev)
					corrected = append(corrected, pos)
					orders[pos][i], orders[pos][i-1] = orders[pos][i-1], orders[pos][i] // swap with prev
					i = -1
				}
			}
		}
	}

	corrected = lib.Unique(corrected)

	sum := 0
	for _, c := range corrected {
		fmt.Printf("corrected %v\n", orders[c])

		l := len(orders[c])
		pos := int(math.Ceil(float64(l / 2)))

		sum += orders[c][pos]
	}

	return sum
}

type IntMaps map[int][]int

func parseRules(section string) (IntMaps, IntMaps) {
	before := make(IntMaps, 0)
	after := make(IntMaps, 0)

	for _, rule := range lib.StrToIntMatrix(section, "|") {
		n0 := rule[0]
		n1 := rule[1]

		{
			r, has := before[n0]

			if has {
				before[n0] = append(r, n1)
			} else {
				before[n0] = []int{n1}
			}
		}

		{
			r, has := after[n1]

			if has {
				after[n1] = append(r, n0)
			} else {
				after[n1] = []int{n0}
			}
		}
	}

	return before, after
}
