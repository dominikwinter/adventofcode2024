package lib

import "slices"

func Remove[T any](slice []T, i int) []T {
	cpy := make([]T, len(slice))

	copy(cpy, slice)

	return slices.Delete(cpy, i, i+1)
}
