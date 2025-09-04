package main

func main() {
}

func subarraySum(nums []int, k int) int {
	hashSum := make(map[int]int, len(nums)+1)
	hashSum[0] = 1
	currSum, counter := 0, 0

	for _, num := range nums {
		currSum += num
		counter += hashSum[currSum-k]
		hashSum[currSum]++
	}

	return counter
}
