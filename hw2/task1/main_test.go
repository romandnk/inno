package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindIntersections(t *testing.T) {
	testCases := []struct {
		name     string
		in       [][]int
		expected []int
	}{
		{
			name: "two non empty slices",
			in: [][]int{
				[]int{1, 2, 3, 2},
				[]int{3, 2},
			},
			expected: []int{2, 3},
		},
		{
			name: "one empty slice",
			in: [][]int{
				[]int{},
			},
			expected: nil,
		},
		{
			name: "one empty slice and two non empty slices",
			in: [][]int{
				[]int{1, 2, 3, 2},
				[]int{3, 2},
				[]int{},
			},
			expected: nil,
		},
		{
			name: "two empty slices",
			in: [][]int{
				[]int{},
				[]int{},
			},
			expected: nil,
		},
		{
			name: "two slices with the same numbers",
			in: [][]int{
				[]int{2, 2, 2, 2},
				[]int{2, 2, 2, 2},
			},
			expected: []int{2},
		},
		{
			name:     "empty input data",
			in:       [][]int{},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := findIntersections(tc.in...)
			require.EqualValues(t, tc.expected, actual)
		})
	}
}
