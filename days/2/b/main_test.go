package main

import (
	"reflect"
	"strings"
	"testing"
)

const INPUT = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

const WANT = 4

func TestRun(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		input := strings.Split(strings.TrimSpace(INPUT), "\n")

		if got := Run(input); !reflect.DeepEqual(got, WANT) {
			t.Errorf("Run() = %v, want %v", got, WANT)
		}
	})
}
