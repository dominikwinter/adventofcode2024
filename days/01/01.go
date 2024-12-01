package days

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	if len(os.Args) < 3 {
		println("Please provide second argument")
		return
	}

	file := os.Args[2]

	in, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	sc := bufio.NewScanner(in)

	lefts := make([]int, 0)
	rights := make([]int, 0)

	for sc.Scan() {
		cells := strings.Split(sc.Text(), "   ")

		left, err := strconv.Atoi(cells[0])
		if err != nil {
			panic(err)
		}

		lefts = append(lefts, left)

		right, err := strconv.Atoi(cells[1])
		if err != nil {
			panic(err)
		}

		rights = append(rights, right)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	dist := 0

	for i := 0; i < len(lefts); i++ {
		dist += Abs(lefts[i] - rights[i])
	}

	println(dist)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
