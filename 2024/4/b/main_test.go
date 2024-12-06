package main

import (
	"reflect"
	"strings"
	"testing"
)

const INPUT = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

const WANT = 9

func TestRun(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		if got := Run(strings.TrimSpace(INPUT)); !reflect.DeepEqual(got, WANT) {
			t.Errorf("Run() = %v, want %v", got, WANT)
		}
	})
}
