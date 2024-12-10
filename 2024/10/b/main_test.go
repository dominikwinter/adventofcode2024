package main

import (
	"reflect"
	"strings"
	"testing"
)

const INPUT = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

const WANT = 81

func TestRun(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		if got := Run(strings.TrimSpace(INPUT)); !reflect.DeepEqual(got, WANT) {
			t.Errorf("Run() = %v, want %v", got, WANT)
		}
	})
}
