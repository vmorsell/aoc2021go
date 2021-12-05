// Code generated by lib/generate_runners.go; DO NOT EDIT

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	name string
	fn   execFn
	want int
}

func Test(t *testing.T) {
	tests := []test{
		{
			name: "part 1",
			fn: part1,
			want: wantPart1Test,
		},
		{
			name: "part 2",
			fn: part2,
			want: wantPart2Test,
		},
	}
	runTests(t, "03.test", parse, tests...)
}

func runTests(t *testing.T, path string, parse parseFn, tests ...test) {
	require.NotNil(t, parse)

	require.NotNil(t, path)
	file, err := os.Open(path)
	require.Nil(t, err)
	defer file.Close()

	input, err := parse(file)
	require.Nil(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.fn(input)
			require.Nil(t, err)
			require.Equal(t, tt.want, res)
		})
	}
}
