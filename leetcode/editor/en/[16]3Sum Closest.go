// Given an integer array nums of length n and an integer target, find three
// integers in nums such that the sum is closest to target.
//
// Return the sum of the three integers.
//
// You may assume that each input would have exactly one solution.
//
// Example 1:
//
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
//
// Example 2:
//
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
//
// Constraints:
//
// 3 <= nums.length <= 500
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
// Related Topics Array Two Pointers Sorting 👍 8897 👎 476
package main

import (
	"math"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	var res int
	diff := 1 << 31
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
			if int(math.Abs(float64(sum-target))) < diff {
				res = sum
				diff = int(math.Abs(float64(sum - target)))
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/31 21:47:32
//Success:
//Runtime:7 ms, faster than 96.34% of Go online submissions.
//Memory Usage:3.2 MB, less than 8.05% of Go online submissions.
