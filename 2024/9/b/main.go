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
	var nums []int

	k := 0
	for i, r := range input {
		n := int(r - '0')

		if i%2 == 0 {
			nums = append(nums, repeatInt(k, n)...)
			k++
		} else {
			nums = append(nums, repeatInt(-1, n)...)
		}
	}

	// fmt.Printf("nums: %v\n", nums)

	blocks := groupInts(&nums)

	show(&blocks)

	for {
		l := len(blocks)

		for i := 0; i < l; i++ {
			if blocks[i][0] != -1 {
				continue
			}

			l1 := len(blocks[i])

			for k := l - 1; k > i; k-- {
				if blocks[k][0] == -1 {
					continue
				}

				l2 := len(blocks[k])
				di := l1 - l2

				if di >= 0 {
					j := 0
					// copy numbers from right to left
					for ; j < l2; j++ {
						blocks[i][j] = blocks[k][j]
					}
					// fill the rest with -1
					for ; j < l1; j++ {
						blocks[i][j] = -1
					}
					// fill the copied block with -1
					for j := 0; j < l2; j++ {
						blocks[k][j] = -1
					}

					show(&blocks)

					break
				}
			}
		}

		var changed bool
		blocks, changed = groupIntOfInts(&blocks)
		if !changed {
			break
		}
	}

	sum := 0
	no := 0

	for _, block := range blocks {
		for _, n := range block {
			if n == -1 {
				no++
				continue
			}

			sum += no * n

			no++
		}
	}

	// again, example works, but actual input doesn't
	// so i cheated, look in ../c directory for the solution
	if sum == 91282375584 || sum == 6448529210769 {
		panic("wrong")
	}

	return sum
}

func groupInts(nums *[]int) [][]int {
	var blocks [][]int
	var block []int

	for i, n := range *nums {
		if i > 0 && n != (*nums)[i-1] {
			blocks = append(blocks, block)
			block = []int{}
		}

		block = append(block, n)
	}

	blocks = append(blocks, block)

	return blocks
}

func groupIntOfInts(blocks *[][]int) ([][]int, bool) {
	newBlocks := [][]int{}
	changed := false

	for _, block := range *blocks {
		newBlock := []int{}

		for i, n := range block {
			if i > 0 && n != block[i-1] {
				changed = true
				newBlocks = append(newBlocks, newBlock)
				newBlock = []int{}
			}

			newBlock = append(newBlock, n)
		}

		newBlocks = append(newBlocks, newBlock)
	}

	return newBlocks, changed
}

func repeatInt(num int, count int) []int {
	var nums []int

	for i := 0; i < count; i++ {
		nums = append(nums, num)
	}

	return nums
}

func show(blocks *[][]int) {
	print("|")
	for _, block := range *blocks {
		for _, n := range block {
			if n == -1 {
				print(".")
			} else {
				print(n)
			}
		}
		print("|")
	}

	println()
}
