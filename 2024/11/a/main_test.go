package main

import (
	"reflect"
	"strings"
	"testing"
)

const INPUT = `
125 17
`

const WANT = 55312

func TestRun(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		if got := Run(strings.TrimSpace(INPUT)); !reflect.DeepEqual(got, WANT) {
			t.Errorf("Run() = %v, want %v", got, WANT)
		}
	})
}
