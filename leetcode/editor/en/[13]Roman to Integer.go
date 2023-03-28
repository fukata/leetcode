// Roman numerals are represented by seven different symbols: I, V, X, L, C, D
// and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
// For example, 2 is written as II in Roman numeral, just two ones added
// together. 12 is written as XII, which is simply X + II. The number 27 is written as
// XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right.
// However, the numeral for four is not IIII. Instead, the number four is written as
// IV. Because the one is before the five we subtract it making four. The same
// principle applies to the number nine, which is written as IX. There are six
// instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
//
// Given a roman numeral, convert it to an integer.
//
// Example 1:
//
// Input: s = "III"
// Output: 3
// Explanation: III = 3.
//
// Example 2:
//
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
//
// Example 3:
//
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
//
// Constraints:
//
// 1 <= s.length <= 15
// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// It is guaranteed that s is a valid roman numeral in the range [1, 3999].
//
// Related Topics Hash Table Math String 👍 9733 👎 567
package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func romanToInt(s string) int {
	roman := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	keys := []string{
		"CM",
		"CD",
		"XC",
		"XL",
		"IX",
		"IV",
		"M",
		"D",
		"C",
		"L",
		"X",
		"V",
		"I",
	}

	result := 0

	for _, k := range keys {
		len1 := len(s)
		if len1 == 0 {
			break
		}

		s = strings.ReplaceAll(s, k, "")
		len2 := len(s)

		diff := len1 - len2
		if diff > 0 {
			n := diff / len(k)
			result += n * roman[k]
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/28 22:37:41
//Success:
//Runtime:20 ms, faster than 7.59% of Go online submissions.
//Memory Usage:5 MB, less than 5.98% of Go online submissions.
