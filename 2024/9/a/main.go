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
			for j := 0; j < n; j++ {
				nums = append(nums, k)
			}
			k++
		} else {
			for j := 0; j < n; j++ {
				nums = append(nums, -1)
			}
		}
	}

	show(&nums)

	k = len(nums) - 1
	for i, v := range nums {
		if i >= k {
			break
		}

		if v == -1 {
			for k > i && nums[k] == -1 {
				k--
			}

			if k > i {
				nums[i], nums[k] = nums[k], nums[i]
				k--
			}

			// show(&nums)
		}
	}

	sum := 0
	for i, v := range nums {
		if v == -1 {
			break
		}

		sum += i * v
	}

	return sum
}

func show(nums *[]int) {
	for _, v := range *nums {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Printf("%v", v)
		}
	}
	fmt.Println()
}
