package main

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
