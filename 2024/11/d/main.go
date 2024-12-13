package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func Run(input string) any {
	stones := lib.SliceToInts(strings.Fields(input))

	stonesMap := blink(stones, 75)

	return sum(stonesMap)
}

func blinkImpl(st map[int]int) map[int]int {
	stones := make(map[int]int)

	for stone, count := range st {
		strStone := strconv.Itoa(stone)
		n := len(strStone)

		if stone == 0 {
			stones[1] += count
		} else if n%2 == 0 {
			stones[lib.ToInt(strStone[:n/2])] += count
			stones[lib.ToInt(strStone[n/2:])] += count
		} else {
			stones[stone*2024] += count
		}
	}
	return stones
}

func blink(stonesI []int, iterations int) map[int]int {
	stones := make(map[int]int)

	for _, st := range stonesI {
		stones[st]++
	}

	for i := 0; i < iterations; i++ {
		stones = blinkImpl(stones)
	}

	return stones
}

func sum(stones map[int]int) int {
	sum := 0

	for _, v := range stones {
		sum += v
	}

	return sum
}
