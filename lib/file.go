package lib

import (
	"bufio"
	"fmt"
)

func Scan(scanner *bufio.Scanner, cb func(line string) error) error {
	for scanner.Scan() {
		if err := cb(scanner.Text()); err != nil {
			return fmt.Errorf("error in callback: %w", err)
		}
	}

	return nil
}
