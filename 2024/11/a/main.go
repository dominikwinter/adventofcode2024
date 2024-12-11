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
	nums := []int{}
	for _, num := range strings.Split(input, " ") {
		nums = append(nums, lib.ToInt(num))
	}

	for step := 0; step < 25; step++ {
		nums2 := []int{}
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				nums2 = append(nums2, 1)
			} else {
				str := strconv.Itoa(nums[i])
				if len(str)%2 == 0 {
					nums2 = append(nums2,
						lib.ToInt(str[:len(str)/2]),
						lib.ToInt(str[len(str)/2:]),
					)
				} else {
					nums2 = append(nums2, nums[i]*2024)
				}
			}
		}

		nums = nums2

		// fmt.Println(nums)
	}

	res := len(nums)

	return res
}
