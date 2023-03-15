// Given a string s, return the longest palindromic substring in s.
//
// Example 1:
//
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
//
// Example 2:
//
// Input: s = "cbbd"
// Output: "bb"
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consist of only digits and English letters.
//
// Related Topics String Dynamic Programming 👍 24280 👎 1418
package main

// 1. 文字列の前後からforループで文字列を見つける。
// 2. 前の値に対応する文字が見つかったら文字列を取り出す。
// 3. 回文かどうかチェックし、文字列長が最大であれば結果変数に格納する。

// leetcode submit region begin(Prohibit modification and deletion)
import "strings"

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func longestPalindrome(s string) string {
	result := ""
	resultLen := 0

	n := len(s)
	for i := 0; i < n; i++ {
		if resultLen > n-i {
			break
		}

		search := s[i : i+1]
		j := n
		for {
			if j < i {
				break
			}

			idx := strings.LastIndex(s[i:j], search)
			if idx == -1 {
				break
			}

			if idx == j {
				j--
			} else {
				j = i + idx
			}

			str := s[i : i+idx+1]
			if resultLen < 1 && len(str) == 1 {
				result = str
				resultLen = 1
				break
			}
			if len(str) > resultLen && str == reverse(str) {
				result = str
				resultLen = len(str)
				break
			}
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
