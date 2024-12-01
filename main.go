package main

import (
	day01 "adventofcode2024/days"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println(`
Usage: go run main.go <day> <command>

Commands:
	test	Run tests
	run	Run the solution (default)
`)

		os.Exit(1)
	}

	day := os.Args[1]
	command := "run"

	if len(os.Args) > 2 {
		command = os.Args[2]
	}

	var file string

	switch command {
	case "test":
		file = "days/" + day + "/example.txt"
	case "run":
		file = "days/" + day + "/input.txt"
	default:
		panic("Invalid command")
	}

	dayFunc := map[string]func(string){
		"1a": day01.Day1A,
		"1b": day01.Day1B,
	}

	if fn, ok := dayFunc[day]; ok {
		fn(file)
	} else {
		panic("Invalid day")
	}
}
