// Given a signed 32-bit integer x, return x with its digits reversed. If
// reversing x causes the value to go outside the signed 32-bit integer range [-2³¹, 2³¹ -
// 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed
// or unsigned).
//
// Example 1:
//
// Input: x = 123
// Output: 321
//
// Example 2:
//
// Input: x = -123
// Output: -321
//
// Example 3:
//
// Input: x = 120
// Output: 21
//
// Constraints:
//
// -2³¹ <= x <= 2³¹ - 1
//
// Related Topics Math 👍 10049 👎 11807
package main

import (
	"math"
	"strconv"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	if -10 < x && x < 10 {
		return x
	}

	min := -2147483648
	max := 2147483647
	result := 0

	xStr := strconv.Itoa(int(math.Abs(float64(x))))
	n := len(xStr)
	for i := n; i > 0; i-- {
		digit, _ := strconv.Atoi(xStr[i-1 : i])
		if digit == 0 {
			continue
		}

		if n > 1 {
			result += digit * int(math.Pow10(i-1))
		} else {
			result += digit
		}

		if result < min || result > max {
			return 0
		}
	}

	if x >= 0 {
		return result
	} else {
		return result * -1
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/17 22:38:24
//Success:
//Runtime:7 ms, faster than 12.01% of Go online submissions.
//Memory Usage:2.1 MB, less than 17.84% of Go online submissions.
