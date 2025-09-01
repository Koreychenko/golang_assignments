# Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


**Example 1:**

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


**Constraints:**

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.

## Benchmarks

There are two approaches: The naive one with O(n^2) time complexity and the expected one with O(n)

You can compare both by running:
```
go test -bench TwoSum -benchmem -run ^$
```

```
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkTwoSumNaive/n=100-12             656857              1559 ns/op              16 B/op          1 allocs/op
BenchmarkTwoSumNaive/n=1k-12                8803            136026 ns/op              16 B/op          1 allocs/op
BenchmarkTwoSumNaive/n=10k-12                952           1223228 ns/op              16 B/op          1 allocs/op
BenchmarkTwoSumNaive/n=100k_end-12         23727             49175 ns/op              16 B/op          1 allocs/op
BenchmarkTwoSumOptimal/n=100-12           429562              3606 ns/op            2737 B/op          3 allocs/op
BenchmarkTwoSumOptimal/n=1k-12             32455             44102 ns/op           41010 B/op          3 allocs/op
BenchmarkTwoSumOptimal/n=10k-12             5636            178834 ns/op          319531 B/op          3 allocs/op
BenchmarkTwoSumOptimal/n=100k_end-12        4995            261288 ns/op         2506797 B/op          3 allocs/op

```

But the tricky part is that The Expected approach is much more memory-hungry :-)
So, in case you have not so many elements, but very limited by memory, the naive approach may be good enough from
the speed perspective as it consumes nothing in terms of memory. 

There is a trick with allocation of memory for the hash table.

```go
hash := make(map[int]int, len(nums))
```

I created one more version without in advance allocation and, in some cases, even though there is a lot of allocations,
it's faster than pre-allocated version.

Why? That's easy! In a real run as the numbers are random there is a high chance that the solution requires less memory
for hash map than len(nums). And the largest len of nums the higher chance it is.    

```
BenchmarkTwoSumOptimalNonAllocated/n=100-12               138330              9023 ns/op            5193 B/op         13 allocs/op
BenchmarkTwoSumOptimalNonAllocated/n=1k-12                 21076             86729 ns/op           43380 B/op         45 allocs/op
BenchmarkTwoSumOptimalNonAllocated/n=10k-12                 3828            384316 ns/op          173171 B/op         98 allocs/op
BenchmarkTwoSumOptimalNonAllocated/n=100k_end-12           13599            113898 ns/op           86463 B/op         61 allocs/op
```

Conclusion for job interviews: 

1. The naive approach is always worse than the optimal one? No, that's wrong. In case you're very limited by memory and the `len(nums)`
is about 200-500 they are almost identical from the speed perspective, but the naive approach is preferable, because it consumes almost no memory.
2. Always pre-allocate memory for hash map? No, that's also not always true. In case of large `len(nums)` and very random distribution of numbers inside
the input slice, non pre-allocated approach may give you higher speed and less memory consumption in the end.
