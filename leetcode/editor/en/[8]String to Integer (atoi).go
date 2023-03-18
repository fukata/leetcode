// Implement the myAtoi(string s) function, which converts a string to a 32-bit
// signed integer (similar to C/C++'s atoi function).
//
// The algorithm for myAtoi(string s) is as follows:
//
// Read in and ignore any leading whitespace.
// Check if the next character (if not already at the end of the string) is '-'
// or '+'. Read this character in if it is either. This determines if the final
// result is negative or positive respectively. Assume the result is positive if
// neither is present.
// Read in next the characters until the next non-digit character or the end of
// the input is reached. The rest of the string is ignored.
// Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If
// no digits were read, then the integer is 0. Change the sign as necessary (from
// step 2).
// If the integer is out of the 32-bit signed integer range [-2³¹, 2³¹ - 1],
// then clamp the integer so that it remains in the range. Specifically, integers
// less than -2³¹ should be clamped to -2³¹, and integers greater than 2³¹ - 1 should
// be clamped to 2³¹ - 1.
// Return the integer as the final result.
//
// Note:
//
// Only the space character ' ' is considered a whitespace character.
// Do not ignore any characters other than the leading whitespace or the rest
// of the string after the digits.
//
// Example 1:
//
// Input: s = "42"
// Output: 42
// Explanation: The underlined characters are what is read in, the caret is the
// current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
//
//	^
//
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
//
//	^
//
// Step 3: "42" ("42" is read in)
//
//	^
//
// The parsed integer is 42.
// Since 42 is in the range [-2³¹, 2³¹ - 1], the final result is 42.
//
// Example 2:
//
// Input: s = "   -42"
// Output: -42
// Explanation:
// Step 1: "   -42" (leading whitespace is read and ignored)
//
//	^
//
// Step 2: "   -42" ('-' is read, so the result should be negative)
//
//	^
//
// Step 3: "   -42" ("42" is read in)
//
//	^
//
// The parsed integer is -42.
// Since -42 is in the range [-2³¹, 2³¹ - 1], the final result is -42.
//
// Example 3:
//
// Input: s = "4193 with words"
// Output: 4193
// Explanation:
// Step 1: "4193 with words" (no characters read because there is no leading
// whitespace)
//
//	^
//
// Step 2: "4193 with words" (no characters read because there is neither a '-'
// nor '+')
//
//	^
//
// Step 3: "4193 with words" ("4193" is read in; reading stops because the next
// character is a non-digit)
//
//	^
//
// The parsed integer is 4193.
// Since 4193 is in the range [-2³¹, 2³¹ - 1], the final result is 4193.
//
// Constraints:
//
// 0 <= s.length <= 200
// s consists of English letters (lower-case and upper-case), digits (0-9), ' ',
// '+', '-', and '.'.
//
// Related Topics String 👍 3047 👎 9248
package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func myAtoi(s string) int {
	result := 0
	negative := false
	parsePhase := false

L:
	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		digit := 0
		switch c {
		case "0":
			parsePhase = true
			if result == 0 {
				continue
			}
			digit = 0
		case "1":
			parsePhase = true
			digit = 1
		case "2":
			parsePhase = true
			digit = 2
		case "3":
			parsePhase = true
			digit = 3
		case "4":
			parsePhase = true
			digit = 4
		case "5":
			parsePhase = true
			digit = 5
		case "6":
			parsePhase = true
			digit = 6
		case "7":
			parsePhase = true
			digit = 7
		case "8":
			parsePhase = true
			digit = 8
		case "9":
			parsePhase = true
			digit = 9
		case " ":
			if parsePhase {
				break L
			}
			continue
		case "-":
			if parsePhase {
				break L
			}
			parsePhase = true
			if result != 0 {
				break L
			}
			negative = true
			continue
		case "+":
			if parsePhase {
				break L
			}
			parsePhase = true
			if result != 0 {
				break L
			}
			negative = false
			continue
		default:
			break L
		}

		result = result*10 + digit

		resultTmp := result
		if negative {
			resultTmp = result * -1
		}
		if resultTmp < math.MinInt32 {
			return math.MinInt32
		}
		if resultTmp > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	if negative {
		result = result * -1
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/18 09:30:27
//Success:
//Runtime:4 ms, faster than 40.81% of Go online submissions.
//Memory Usage:2.2 MB, less than 59.79% of Go online submissions.
