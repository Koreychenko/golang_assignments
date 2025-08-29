package main

func main() {
}

func twoSum(nums []int, target int) []int {
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
