package lib

import (
	"cmp"
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

	for range steps {
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

	for range steps {
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

	for i := range height {
		for k := range width {
			output[i+k] += input[i][k : k+1]
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

func Unique[S ~[]E, E cmp.Ordered](list S) S {
	slices.Sort(list)
	return slices.Compact(list)
}

func GetColumn[T any](matrix [][]T, col int) []T {
	res := make([]T, len(matrix))

	for i, row := range matrix {
		res[i] = row[col]
	}

	return res
}

func GetSorted[S ~[]E, E cmp.Ordered](list S) S {
	asc := make(S, len(list))

	copy(asc, list)

	slices.Sort(asc)

	return asc
}

func GetReverse[S ~[]E, E cmp.Ordered](list S) S {
	rev := make(S, len(list))

	copy(rev, list)

	slices.Reverse(rev)

	return rev
}

func UniqueAdd[S ~[]E, E comparable](s *S, v E) {
	if !slices.Contains(*s, v) {
		*s = append(*s, v)
	}
}

func Map[I any, O any](list []I, f func(I) O) []O {
	res := make([]O, len(list))

	for i, s := range list {
		res[i] = f(s)
	}

	return res
}
