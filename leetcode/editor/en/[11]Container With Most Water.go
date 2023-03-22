// You are given an integer array height of length n. There are n vertical lines
// drawn such that the two endpoints of the iᵗʰ line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the
// container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.
//
// Example 1:
//
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,
// 3,7]. In this case, the max area of water (blue section) the container can
// contain is 49.
//
// Example 2:
//
// Input: height = [1,1]
// Output: 1
//
// Constraints:
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
//
// Related Topics Array Two Pointers Greedy 👍 23270 👎 1241
package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	result := 0

	start := 0
	end := len(height) - 1
	for start < end {
		currentArea := end - start
		if height[start] < height[end] {
			currentArea = currentArea * height[start]
			start++
		} else {
			currentArea = currentArea * height[end]
			end--
		}

		if currentArea > result {
			result = currentArea
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/22 23:53:11
//Success:
//Runtime:90 ms, faster than 23.95% of Go online submissions.
//Memory Usage:10.5 MB, less than 5.24% of Go online submissions.
