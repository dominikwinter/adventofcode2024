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

func Scale(x, in_min, in_max, out_min, out_max float64) float64 {
	return ((x-in_min)/(in_max-in_min))*(out_max-out_min) + out_min
}

func Max[T int | int32 | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T int | int32 | int64 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}

	return b
}
