package main

import (
	"reflect"
	"strings"
	"testing"
)

const INPUT = `
3   4
4   3
2   5
1   3
3   9
3   3
`

const WANT = 11

func TestRun(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		input := strings.Split(strings.TrimSpace(INPUT), "\n")

		if got := Run(input); !reflect.DeepEqual(got, WANT) {
			t.Errorf("Run() = %v, want %v", got, WANT)
		}
	})
}
