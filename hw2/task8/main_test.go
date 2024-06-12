package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsEqualArrays(t *testing.T) {
	tests := []struct {
		name          string
		first, second []int
		expected      bool
	}{
		{
			name:     "equal",
			first:    []int{3, 4, 2, 9, 1, 5, 8, 2, 6},
			second:   []int{9, 8, 2, 6, 5, 3, 4, 2, 1},
			expected: true,
		},
		{
			name:     "equal",
			first:    []int{2, 2, 2, 9, 1, 5, 2, 2, 6},
			second:   []int{9, 1, 2, 2, 5, 2, 6, 2, 2},
			expected: true,
		},
		{
			name:     "empty",
			first:    []int{},
			second:   []int{},
			expected: true,
		},
		{
			name:     "with similar values",
			first:    []int{1, 1, 1, 1, 1},
			second:   []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "with similar values",
			first:    []int{5, 2, 2, 4, 6, 6},
			second:   []int{6, 5, 2, 2, 4, 6},
			expected: true,
		},
		{
			name:     "not equal length",
			first:    []int{1, 2, 3, 4, 2, 2},
			second:   []int{1, 2, 3},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := IsEqualArrays(test.first, test.second)
			require.Equal(t, test.expected, actual)
		})
	}
}
