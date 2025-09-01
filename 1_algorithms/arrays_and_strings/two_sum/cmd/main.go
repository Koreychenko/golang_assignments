package main

import (
	"cmp"
	"slices"
)

func main() {
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))

	for i, val := range nums {
		i1, ok := hash[target-val]

		if ok {
			return []int{i1, i}
		} else {
			hash[val] = i
		}
	}

	return []int{0, 0}
}

func twoSumTwoPointers(nums []int, target int) []int {
	type indexPair struct {
		value         int
		originalIndex int
	}

	newNums := make([]indexPair, len(nums))

	for i, val := range nums {
		newNums[i] = indexPair{val, i}
	}

	slices.SortFunc(newNums, func(a, b indexPair) int {
		return cmp.Compare(a.value, b.value)
	})

	pointer1 := 0
	pointer2 := len(nums) - 1

	for pointer1 < pointer2 {
		if newNums[pointer1].value+newNums[pointer2].value < target {
			pointer1++
		} else if newNums[pointer1].value+newNums[pointer2].value > target {
			pointer2--
		} else {
			return []int{newNums[pointer1].originalIndex, newNums[pointer2].originalIndex}
		}
	}

	return []int{0, 0}
}

func twoSumNonAllocated(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, val := range nums {
		i1, ok := hash[target-val]

		if ok {
			return []int{i1, i}
		} else {
			hash[val] = i
		}
	}

	return []int{0, 0}
}

func twoSumNaive(nums []int, target int) []int {
	for i1 := 0; i1 < len(nums); i1++ {
		for i2 := i1 + 1; i2 < len(nums); i2++ {
			if nums[i1]+nums[i2] == target {
				return []int{i1, i2}
			}
		}
	}

	return []int{0, 0}
}

func twoSumMappingTable(nums []int, target int) []int {
	minV, maxV := nums[0], nums[0]
	for _, v := range nums {
		if v < minV {
			minV = v
		}
		if v > maxV {
			maxV = v
		}
	}

	if maxV-minV > 1_000_000_000 {
		return []int{0, 0}
	}

	offset := -minV
	size := maxV - minV + 1

	// store first index seen for each value; -1 means unseen
	first := make([]int, size)

	for idx, v := range nums {
		need := target - v
		ni := need + offset
		if 0 <= ni && ni < size && first[ni] != 0 {
			return []int{first[ni] - 1, idx}
		}
		vi := v + offset
		if first[vi] == 0 {
			first[vi] = idx + 1
		}
	}
	return []int{0, 0}
}
