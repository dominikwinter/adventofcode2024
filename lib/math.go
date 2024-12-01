package lib

func Abs[T int | int32 | int64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}

	return x
}
