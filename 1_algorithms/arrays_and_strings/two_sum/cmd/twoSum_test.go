package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		nums     []int
		target   int
		expected []int
	}

	testcases := []testCase{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.expected, twoSum(tc.nums, tc.target))
	}
}
