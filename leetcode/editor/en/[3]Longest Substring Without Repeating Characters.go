// Given a string s, find the length of the longest substring without repeating
// characters.
//
// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a
// substring.
//
// Constraints:
//
// 0 <= s.length <= 5 * 10⁴
// s consists of English letters, digits, symbols and spaces.
//
// Related Topics Hash Table String Sliding Window 👍 32635 👎 1425
package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	result := 0
	word := ""
	wordLen := 0

	n := len(s)
	for i := 0; i < n; i++ {
		char := s[i : i+1]

		idx := strings.Index(word, char)
		if idx > -1 {
			if idx+1 == len(word) {
				word = ""
				wordLen = 0
			} else {
				word = word[idx+1:]
				wordLen = len(word)
			}
		}

		word += char
		wordLen += 1

		if wordLen > result {
			result = wordLen
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
