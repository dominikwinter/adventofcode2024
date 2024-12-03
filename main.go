package main

import (
	day01 "adventofcode2024/days"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println(`
Usage: go run main.go <day> <file>

Commands:
	test	Run tests
	file	Input file
`)

		os.Exit(1)
	}

	day := os.Args[1]
	file := os.Args[2]

	dayFunc := map[string]func(*bufio.Scanner){
		"1a": day01.Day1A, "1b": day01.Day1B,
		"2a": day01.Day2A, "2b": day01.Day2B,
		"3a": day01.Day3A, "3b": day01.Day3B,
	}

	if fn, ok := dayFunc[day]; ok {
		in, err := os.Open(file)

		if err != nil {
			panic(fmt.Errorf("error opening file: %w", err))
		}

		defer in.Close()

		scanner := bufio.NewScanner(in)
		fn(scanner)
	} else {
		panic("Invalid day")
	}
}
