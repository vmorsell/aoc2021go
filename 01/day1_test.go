package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	in := `199
200
208
210
200
207
240
269
260
263`

	input, err := input(strings.NewReader(in))
	require.Nil(t, err)

	t.Run("part 1", func(t *testing.T) {
		want := 7
		got := part1(input)
		require.Equal(t, got, want)
	})

	t.Run("part 2", func(t *testing.T) {
		want := 5
		got := part2(input)
		require.Equal(t, got, want)
	})
}
