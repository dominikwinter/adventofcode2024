package main

import (
	day01 "adventofcode2024/days/01"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Please provide an argument")
		return
	}

	day := os.Args[1]

	switch day {
	case "1":
		day01.Run()
	default:
		panic("Invalid day")
	}
}
