package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vmorsell/aoc2021go/lib"
)

func main() {
	f, err := os.Open("day2.in")
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	input, err := input(f)
	if err != nil {
		log.Fatalf("input: %v", err)
	}
	if err != nil {
		log.Fatalf("read lines: %v", err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type move struct {
	direction string
	steps     int
}

func input(r io.Reader) ([]move, error) {
	var out []move
	err := lib.ReadLines(r, func(s string) error {
		args := strings.Split(s, " ")
		if len(args) != 2 {
			return fmt.Errorf("got %d args want 2", len(args))
		}
		steps, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("atoi: %w", err)
		}
		out = append(out, move{
			direction: args[0],
			steps:     steps,
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read lines: %w", err)
	}
	return out, nil
}

func part1(moves []move) int {
	horiz := 0
	depth := 0

	for _, m := range moves {
		switch m.direction {
		case "forward":
			horiz += m.steps
		case "down":
			depth += m.steps
		case "up":
			depth -= m.steps
		}
	}
	return horiz * depth
}

func part2(moves []move) int {
	horiz := 0
	depth := 0
	aim := 0

	for _, m := range moves {
		switch m.direction {
		case "forward":
			horiz += m.steps
			depth += aim * m.steps
		case "down":
			aim += m.steps
		case "up":
			aim -= m.steps
		}
	}
	return horiz * depth
}
