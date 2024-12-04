package lib

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"12345", "54321"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestAllIndexes(t *testing.T) {
	tests := []struct {
		str      string
		substr   string
		expected []int
	}{
		{"hello world", "o", []int{4, 7}},
		{"hello world", "l", []int{2, 3, 9}},
		{"hello world", "world", []int{6}},
		{"hello world", "hello", []int{0}},
		{"hello world", "x", []int{}},
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
	}

	for _, test := range tests {
		result := AllIndexes(test.str, test.substr)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("AllIndexes(%q, %q) = %v; expected %v", test.str, test.substr, result, test.expected)
		}
	}
}
