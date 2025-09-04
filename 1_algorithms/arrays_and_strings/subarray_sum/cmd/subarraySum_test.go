package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected int
}

func getTestCases() []testCase {
	return []testCase{
		{
			[]int{1, 1, 1},
			2,
			2,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
	}
}

func TestSubarraySum(t *testing.T) {
	for _, tc := range getTestCases() {
		assert.Equal(t, tc.expected, subarraySum(tc.nums, tc.target))
	}
}
