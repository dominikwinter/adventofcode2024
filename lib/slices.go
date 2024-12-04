package lib

import (
	"fmt"
	"slices"
	"strconv"
)

func Remove[T any](slice []T, i int) []T {
	l := len(slice)

	if i < 0 || i >= l {
		return slice
	}

	cpy := make([]T, l)

	copy(cpy, slice)

	return slices.Delete(cpy, i, i+1)
}

func Transpose[T any](matrix [][]T, steps int) [][]T {
	steps = steps % 4

	if steps == 0 {
		return matrix
	}

	for step := 0; step < steps; step++ {
		n := len(matrix)
		m := len(matrix[0])

		transposed := make([][]T, m)

		for i := range transposed {
			transposed[i] = make([]T, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				transposed[j][n-1-i] = matrix[i][j]
			}
		}

		matrix = transposed
	}

	return matrix
}

func TransposeStrings(matrix []string, steps int) []string {
	steps = steps % 4

	if steps == 0 {
		return matrix
	}

	for step := 0; step < steps; step++ {
		n := len(matrix)
		m := len(matrix[0])

		transposed := make([][]byte, m)

		for i := range transposed {
			transposed[i] = make([]byte, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				transposed[j][n-1-i] = matrix[i][j]
			}
		}

		newMatrix := make([]string, m)
		for i := range transposed {
			newMatrix[i] = string(transposed[i])
		}

		matrix = newMatrix
	}

	return matrix
}

func TransposeDiagonalStrings(input []string) []string {
	height := len(input)

	if height == 0 {
		return input
	}

	width := len(input[0])

	output := make([]string, height+width-1)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			output[i+j] += input[i][j : j+1]
		}
	}

	return output
}

func SliceToInts(list []string) []int {
	res := make([]int, len(list))

	for i, s := range list {
		num, err := strconv.Atoi(s)

		if err != nil {
			fmt.Printf("error parsing int: %v\n", err)
		}

		res[i] = num
	}

	return res
}
