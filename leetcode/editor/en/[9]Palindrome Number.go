// Given an integer x, return true if x is a palindrome, and false otherwise.
//
// Example 1:
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
//
// Example 2:
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it
// becomes 121-. Therefore it is not a palindrome.
//
// Example 3:
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
// Constraints:
//
// -2³¹ <= x <= 2³¹ - 1
//
// Follow up: Could you solve it without converting the integer to a string?
//
// Related Topics Math 👍 9054 👎 2452
package main

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var digits []int
	for ; x > 0; x /= 10 {
		digit := x % 10
		digits = append(digits, digit)
	}

	n := len(digits)
	for i := 0; i < n; i++ {
		if i == n-i {
			break
		}
		if i > n-1 {
			break
		}

		if digits[i] != digits[n-i-1] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/19 09:23:31
//Success:
//Runtime:24 ms, faster than 56.11% of Go online submissions.
//Memory Usage:6.7 MB, less than 11.14% of Go online submissions.
