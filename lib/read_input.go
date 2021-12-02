package lib

import (
	"bufio"
	"io"
)

// ReadLines reads a file line by line and executes the
// provided function for each line.
func ReadLines(reader io.Reader, fn func(s string) error) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		v := scanner.Text()
		fn(v)
	}
	return nil
}
