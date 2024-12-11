package lib

import "strings"

func Reverse(str string) string {
	result := ""

	for _, v := range str {
		result = string(v) + result
	}

	return result
}

func AllIndexes(str string, sub string) []int {
	result := []int{}
	strLen := len(str)
	subLen := len(sub)

	for i := 0; i < strLen; i++ {
		if str[i] == sub[0] {
			if i+subLen <= strLen && str[i:i+subLen] == sub {
				result = append(result, i)
			}
		}
	}

	return result
}

func SplitLines(str string) []string {
	return strings.Split(str, "\n")
}

func StrToStrMatrix(str string, sep string) [][]string {
	var matrix [][]string

	for _, line := range SplitLines(str) {
		matrix = append(matrix, strings.Split(line, sep))
	}

	return matrix
}

func StrToRuneMatrix(str string, sep string) [][]rune {
	var matrix [][]rune

	for _, line := range SplitLines(str) {
		matrix = append(matrix, []rune(line))
	}

	return matrix
}

func StrToIntMatrix(str string, sep string) [][]int {
	var matrix [][]int

	for _, line := range SplitLines(str) {
		matrix = append(matrix, SliceToInts(strings.Split(line, sep)))
	}

	return matrix
}

func GroupStr(s string) []string {
	var list []string
	current := ""

	for i, r := range s {
		if i > 0 && r != rune(s[i-1]) {
			list = append(list, current)
			current = ""
		}

		current += string(r)
	}
	list = append(list, current)

	return list
}
