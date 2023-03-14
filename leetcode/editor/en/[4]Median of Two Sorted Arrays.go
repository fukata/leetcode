// Given two sorted arrays nums1 and nums2 of size m and n respectively, return
// the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
//
// Example 1:
//
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
//
// Example 2:
//
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
// Constraints:
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
// Related Topics Array Binary Search Divide and Conquer 👍 22499 👎 2526
package main

// leetcode submit region begin(Prohibit modification and deletion)
import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArray := append(nums1, nums2...)
	sort.Ints(mergedArray)

	n := len(mergedArray)
	if n%2 == 0 {
		return float64(mergedArray[n/2-1]+mergedArray[n/2]) / 2
	} else {
		return float64(mergedArray[n/2])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
