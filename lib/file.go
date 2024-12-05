package lib

import (
	"bufio"
	"io"
	"strings"
)

func Read(r io.Reader) string {
	scanner := bufio.NewScanner(r)

	input := ""
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	return strings.TrimSpace(input)
}
