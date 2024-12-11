package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	lib "github.com/dominikwinter/adventofcode2024/lib"
)

func main() {
	start := time.Now()
	fmt.Printf("%v\n", Run(lib.Read(os.Stdin)))
	fmt.Printf("\nTime: %v\n", time.Since(start))
}

// very ugly solution, translated from py***
func Run(input string) any {
	files := make(map[int][2]int)
	blanks := [][2]int{}
	fid := 0
	pos := 0

	for i, char := range input {
		x, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			if x == 0 {
				panic("unexpected x=0 for file")
			}
			files[fid] = [2]int{pos, x}
			fid++
		} else {
			if x != 0 {
				blanks = append(blanks, [2]int{pos, x})
			}
		}
		pos += x
	}

	for fid > 0 {
		fid--
		pos, size := files[fid][0], files[fid][1]
		for i, blank := range blanks {
			start, length := blank[0], blank[1]
			if start >= pos {
				blanks = blanks[:i]
				break
			}
			if size <= length {
				files[fid] = [2]int{start, size}
				if size == length {
					blanks = append(blanks[:i], blanks[i+1:]...)
				} else {
					blanks[i] = [2]int{start + size, length - size}
				}
				break
			}
		}
	}

	total := 0
	for fid, file := range files {
		pos, size := file[0], file[1]
		for x := pos; x < pos+size; x++ {
			total += fid * x
		}
	}

	return total
}
