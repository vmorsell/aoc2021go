package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/vmorsell/aoc2021go/lib"
)

func main() {
	f, err := os.Open("day1.in")
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	input, err := input(f)
	if err != nil {
		log.Fatalf("input: %v", err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func input(r io.Reader) ([]int, error) {
	var out []int
	err := lib.ReadLines(r, func(s string) error {
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("atoi: %w", err)
		}
		out = append(out, n)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read int lines: %w", err)
	}
	return out, nil
}

func part1(input []int) int {
	res := 0
	for i := range input {
		if i == 0 {
			continue
		}
		if input[i] > input[i-1] {
			res++
		}
	}
	return res
}

func part2(input []int) int {
	res := 0
	// Two numbers are always common for the floating sums.
	// We only need to need to compare i-1 and i+2.
	// Obviously, this also means we don't iterate over the
	// last two values in the list.
	for i := range input[:len(input)-2] {
		if i == 0 {
			continue
		}
		if input[i+2] > input[i-1] {
			res++
		}
	}
	return res
}
