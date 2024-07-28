package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvalSequence(t *testing.T) {
	type args struct {
		mtx [][]int
		ua  []int
	}

	// 0 -(2)- 1 -(1)- 3
	// |         \
	//(3)         (1)
	// |            \
	// 2             4
	mtx1 := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	tests := []struct {
		name        string
		args        args
		want        int
		expectedErr string
	}{
		{
			name: "mtx 5 verticals 100%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 0, 2},
			},
			want: 100,
		},
		{
			name: "mtx 5 verticals 0%",
			args: args{
				mtx: mtx1,
				ua:  []int{},
			},
			want: 0,
		},
		{
			name: "mtx 5 verticals 50%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 0},
			},
			want: 50,
		},
		{
			name: "mtx 5 verticals 33%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 3},
			},
			want: 33,
		},
		{
			name: "mtx 5 verticals 33%",
			args: args{
				mtx: mtx1,
				ua:  []int{0, 2},
			},
			want: 50,
		},
		{
			name: "matrix has loop",
			args: args{
				mtx: [][]int{
					{1, 2, 3, 0, 0},
					{2, 0, 0, 1, 1},
					{3, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0},
				},
				ua: []int{4, 1, 0},
			},
			want:        0,
			expectedErr: "matrix has loop",
		},
		{
			name: "matrix is empty",
			args: args{
				mtx: [][]int{},
				ua:  []int{4, 1, 0},
			},
			want:        0,
			expectedErr: "matrix is empty",
		},
		{
			name: "matrix is not symmetric",
			args: args{
				mtx: [][]int{
					{0, 2, 3, 0, 0},
					{2, 0, 0, 1, 1},
					{3, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
				},
				ua: []int{4, 1, 0},
			},
			want:        0,
			expectedErr: "matrix is not square",
		},
		{
			name: "matrix is not symmetric",
			args: args{
				mtx: [][]int{
					{0, 2, 3, 0, 0},
					{2, 0, 0, 1, 1},
					{3, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
				},
				ua: []int{4, 1, 0},
			},
			want:        0,
			expectedErr: "matrix is not symmetric",
		},
		{
			name: "duplicated answer",
			args: args{
				mtx: [][]int{
					{0, 2, 3, 0, 0},
					{2, 0, 0, 1, 1},
					{3, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0},
				},
				ua: []int{4, 1, 0, 1},
			},
			want:        0,
			expectedErr: "answer is duplicated: 1",
		},
		{
			name: "negative answer",
			args: args{
				mtx: [][]int{
					{0, 2, 3, 0, 0},
					{2, 0, 0, 1, 1},
					{3, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0},
				},
				ua: []int{4, 1, 0, -1},
			},
			want:        0,
			expectedErr: "invalid answer: -1",
		},
		{
			name: "answer is out of available answers",
			args: args{
				mtx: [][]int{
					{0, 2, 3, 0, 0},
					{2, 0, 0, 1, 1},
					{3, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 1, 0, 0, 0},
				},
				ua: []int{4, 1, 6},
			},
			want:        0,
			expectedErr: "invalid answer: 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalSequence(tt.args.mtx, tt.args.ua)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
