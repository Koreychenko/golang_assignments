package main

import (
	"math/rand"
	"testing"
)

var benchSink []int // prevents compiler from optimizing results away

// makeInput builds a slice of size n and plants a guaranteed solution.
// If solutionAtEnd is true, the matching pair is at positions 0 and n-1
// (stressful for naive); otherwise it's somewhere in the middle.
func makeInput(n int, solutionAtEnd bool) (nums []int, target int) {
	r := rand.New(rand.NewSource(42))
	nums = make([]int, n)
	for i := range nums {
		nums[i] = r.Intn(1_000_000)
	}
	i1, i2 := 0, n-1
	if !solutionAtEnd && n >= 2 {
		i1 = n / 3
		i2 = 2 * n / 3
	}
	target = nums[i1] + nums[i2]
	return
}

func runBench(b *testing.B, n int, solutionAtEnd bool, fn func([]int, int) []int) {
	// Prebuild a few input variants so generation isn’t timed.
	const variants = 8
	inputs := make([][]int, variants)
	targets := make([]int, variants)
	for v := 0; v < variants; v++ {
		nums, tgt := makeInput(n, solutionAtEnd)
		inputs[v] = nums
		targets[v] = tgt
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nums := inputs[i%variants]
		tgt := targets[i%variants]
		benchSink = fn(nums, tgt)
	}
}

func BenchmarkTwoSumNaive(b *testing.B) {
	cases := []struct {
		name          string
		n             int
		solutionAtEnd bool
	}{
		{"n=100", 100, false},
		{"n=1k", 1000, false},
		{"n=10k", 10_000, false},
		{"n=100k_end", 100_000, true}, // worst-ish for naive
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			runBench(b, c.n, c.solutionAtEnd, twoSumNaive)
		})
	}
}

func BenchmarkTwoSumOptimal(b *testing.B) {
	cases := []struct {
		name          string
		n             int
		solutionAtEnd bool
	}{
		{"n=100", 100, false},
		{"n=1k", 1000, false},
		{"n=10k", 10_000, false},
		{"n=100k_end", 100_000, true},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			runBench(b, c.n, c.solutionAtEnd, twoSum)
		})
	}
}
