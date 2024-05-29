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
			first:    []int{1, 2, 3},
			second:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "equal not equal length",
			first:    []int{1, 2, 3, 4, 2, 2},
			second:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "equal not equal length",
			first:    []int{2, 3},
			second:   []int{2, 2, 2, 2},
			expected: true,
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
