// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
// of rows like this: (you may want to display this pattern in a fixed font for
// better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R
//
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a
// number of rows:
//
// string convert(string s, int numRows);
//
// Example 1:
//
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
//
// Example 2:
//
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// Example 3:
//
// Input: s = "A", numRows = 1
// Output: "A"
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000
//
// Related Topics String 👍 5894 👎 11893
package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	// 各行の文字
	result := make([]string, numRows, numRows)
	direction := 1 // 進む方向
	line := 0      // 現在の行
	for i := 0; i < len(s); i++ {
		result[line] += s[i : i+1]

		// 上下両端に来たら進む方向を反転させる
		nextLine := line + direction
		if nextLine >= numRows {
			direction = -1
		} else if nextLine < 0 {
			direction = 1
		}

		line += direction
	}

	return strings.Join(result, "")
}

//leetcode submit region end(Prohibit modification and deletion)
