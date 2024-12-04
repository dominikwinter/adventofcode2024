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
