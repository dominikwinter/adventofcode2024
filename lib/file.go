package lib

import (
	"bufio"
	"fmt"
	"os"
)

func File(file string, cb func(line string) error) error {
	in, err := os.Open(file)

	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer in.Close()

	sc := bufio.NewScanner(in)

	for sc.Scan() {
		if err := cb(sc.Text()); err != nil {
			return fmt.Errorf("error in callback: %w", err)
		}
	}

	return nil
}
