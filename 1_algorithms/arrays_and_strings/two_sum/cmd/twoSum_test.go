package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

func getTestCases() []testCase {
	return []testCase{
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
}

func TestTwoSum(t *testing.T) {
	for _, tc := range getTestCases() {
		assert.Equal(t, tc.expected, twoSum(tc.nums, tc.target))
	}
}

func TestTwoSumNaive(t *testing.T) {
	for _, tc := range getTestCases() {
		assert.Equal(t, tc.expected, twoSumNaive(tc.nums, tc.target))
	}
}
