package lib

import (
	"fmt"
	"strconv"
)

func Abs[T int | int32 | int64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func ToInt(s string) int {
	num, err := strconv.Atoi(s)

	if err != nil {
		fmt.Printf("error parsing int: %v\n", err)
	}

	return num
}
