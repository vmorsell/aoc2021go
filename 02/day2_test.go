package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay2(t *testing.T) {
	in := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	input, err := input(strings.NewReader(in))
	require.Nil(t, err)

	t.Run("part 1", func(t *testing.T) {
		want := 150
		got := part1(input)
		require.Equal(t, want, got)
	})

	t.Run("part 2", func(t *testing.T) {
		want := 900
		got := part2(input)
		require.Equal(t, want, got)
	})
}
