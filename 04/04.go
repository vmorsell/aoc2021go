package main

//go:generate go run ../lib/genrunners 04

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	wantPart1Test int = 4512
	wantPart2Test int = 1924
)

type data struct {
	draws  []int
	boards []*board
}

const boardSize = 5

type board struct {
	nums     []int
	unmarked map[int]struct{}
	hasWon   bool
}

func newBoard() board {
	return board{
		unmarked: make(map[int]struct{}, boardSize*boardSize),
	}
}

func (b *board) mark(n int) bool {
	for _, num := range b.nums {
		if num == n {
			delete(b.unmarked, n)
			return true
		}
	}
	return false
}

func (b *board) won() bool {
	// Check horizontally
nextRow:
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			n := b.nums[i*boardSize+j]
			if _, ok := b.unmarked[n]; ok {
				continue nextRow
			}
		}
		return true
	}

	// Check vertically
nextCol:
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			n := b.nums[i+j*boardSize]
			if _, ok := b.unmarked[n]; ok {
				continue nextCol
			}
		}
		return true
	}

	return false
}

func (b *board) score(draw int) int {
	s := 0
	for k := range b.unmarked {
		s += k
	}
	return s * draw
}

func parse(r io.Reader) (data, error) {
	scanner := bufio.NewScanner(r)

	// Scan numbers draws
	var draws []int
	scanner.Scan()
	strNumbers := strings.Split(scanner.Text(), ",")
	for _, s := range strNumbers {
		n, err := strconv.Atoi(s)
		if err != nil {
			return data{}, fmt.Errorf("atoi: %w", err)
		}
		draws = append(draws, n)
	}

	// Skip empty line
	scanner.Scan()

	// Scan boards
	var boards []*board
	for scanner.Scan() {
		b, err := parseBoard(scanner)
		if err != nil {
			return data{}, fmt.Errorf("scan board: %w", err)
		}
		boards = append(boards, &b)
	}

	return data{
		draws:  draws,
		boards: boards,
	}, nil
}

func parseBoard(s *bufio.Scanner) (board, error) {
	b := newBoard()
	for i := 0; i < boardSize; i++ {
		strNums := strings.Split(s.Text(), " ")
		for _, s := range strNums {
			if s == "" {
				continue
			}

			n, err := strconv.Atoi(s)
			if err != nil {
				return board{}, fmt.Errorf("atoi: %w", err)
			}
			b.nums = append(b.nums, n)
			b.unmarked[n] = struct{}{}
		}
		s.Scan()
	}
	return b, nil
}

func part1(input data) (int, error) {
	for i, d := range input.draws {
		for _, b := range input.boards {
			found := b.mark(d)
			if !found {
				continue
			}

			if i < boardSize {
				continue
			}

			if b.won() {
				return b.score(d), nil
			}
		}
	}
	return 0, fmt.Errorf("no board won :(")
}

func part2(input data) (int, error) {
	boardsWon := 0
	for i, d := range input.draws {
		for _, b := range input.boards {
			if b.hasWon {
				continue
			}

			found := b.mark(d)
			if !found {
				continue
			}

			if i < boardSize {
				continue
			}

			if b.won() {
				b.hasWon = true
				boardsWon++
				if boardsWon == len(input.boards) {
					return b.score(d), nil
				}
			}
		}
	}
	return 0, fmt.Errorf("no board won :(")
}
