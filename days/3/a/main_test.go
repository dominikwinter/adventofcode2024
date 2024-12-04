package main

import (
	"reflect"
	"strings"
	"testing"
)

const INPUT = `
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`

const WANT = 161

func TestRun(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		input := strings.Split(strings.TrimSpace(INPUT), "\n")

		if got := Run(input); !reflect.DeepEqual(got, WANT) {
			t.Errorf("Run() = %v, want %v", got, WANT)
		}
	})
}
