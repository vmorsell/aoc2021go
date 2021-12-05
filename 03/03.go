package main

//go:generate go run ../lib/genrunners 03

import (
	"fmt"
	"io"
	"strconv"

	"github.com/vmorsell/aoc2021go/lib"
)

var (
	wantPart1Test int = 198
	wantPart2Test int = 230
)

func parse(r io.Reader) ([][]bool, error) {
	var out [][]bool
	err := lib.ReadLines(r, func(s string) error {
		var row []bool
		for i := 0; i < len(s); i++ {
			var bit bool
			if s[i:i+1] == "1" {
				bit = true
			}
			row = append(row, bit)
		}
		out = append(out, row)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read lines: %w", err)
	}
	return out, nil
}

func part1(input [][]bool) (int, error) {
	g, err := gamma(input)
	if err != nil {
		return 0, fmt.Errorf("gamma: %w", err)
	}

	e, err := epsilon(input)
	if err != nil {
		return 0, fmt.Errorf("epsilon: %w", err)
	}
	return g * e, nil
}

func gamma(input [][]bool) (int, error) {
	params := common(input, true)

	res, err := binToDec(params)
	if err != nil {
		return 0, fmt.Errorf("bin to dec: %w", err)
	}
	return res, nil
}

func epsilon(input [][]bool) (int, error) {
	params := common(input, false)

	res, err := binToDec(params)
	if err != nil {
		return 0, fmt.Errorf("bin to dec: %w", err)
	}
	return res, nil
}

func common(input [][]bool, useMostCommon bool) []bool {
	var out []bool
	for c := 0; c < len(input[0]); c++ {
		res := 0
		for r := 0; r < len(input); r++ {
			if input[r][c] {
				res++
			} else {
				res--
			}
		}

		var bit bool
		if useMostCommon {
			bit = res > 0
		} else {
			bit = res < 0
		}
		out = append(out, bit)
	}
	return out
}

func binToDec(binary []bool) (int, error) {
	s := ""
	for _, b := range binary {
		v := "0"
		if b {
			v = "1"
		}
		s += v
	}
	dec, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
	}
	return int(dec), err
}

func part2(input [][]bool) (int, error) {
	o, err := oxygen(input)
	if err != nil {
		return 0, fmt.Errorf("oxygen: %w", err)
	}
	c, err := co2(input)
	if err != nil {
		return 0, fmt.Errorf("co2: %w", err)
	}
	return o * c, nil
}

func oxygen(input [][]bool) (int, error) {
	params := search(input, true)

	res, err := binToDec(params)
	if err != nil {
		return 0, fmt.Errorf("bin to dec: %w", err)
	}
	return res, nil
}

func co2(input [][]bool) (int, error) {
	params := search(input, false)

	res, err := binToDec(params)
	if err != nil {
		return 0, fmt.Errorf("bin to dec: %w", err)
	}
	return res, nil
}

func search(input [][]bool, keepMostCommon bool) []bool {
	candidates := make(map[int][]bool, len(input))
	for i, v := range input {
		candidates[i] = v
	}

	for i := 0; i < len(input[0]); i++ {
		var ones []int
		var zeros []int

		for k, v := range candidates {
			if v[i] {
				ones = append(ones, k)
			} else {
				zeros = append(zeros, k)
			}
		}

		drop := ones
		if keepMostCommon {
			if len(ones) == len(zeros) {
				drop = zeros
			}
			if len(ones) > len(zeros) {
				drop = zeros
			}
		} else {
			if len(ones) < len(zeros) {
				drop = zeros
			}
		}
		for _, k := range drop {
			delete(candidates, k)
		}

		if len(candidates) == 1 {
			for _, v := range candidates {
				return v
			}
		}
	}
	return nil
}
