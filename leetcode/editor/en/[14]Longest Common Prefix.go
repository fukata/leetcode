// Write a function to find the longest common prefix string amongst an array of
// strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// Constraints:
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.
//
// Related Topics String Trie 👍 13205 👎 3831
package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		for strings.Index(str, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

//leetcode submit region end(Prohibit modification and deletion)

//> 2023/03/29 16:53:25
//Success:
//Runtime:0 ms, faster than 100.00% of Go online submissions.
//Memory Usage:2.3 MB, less than 46.91% of Go online submissions.
//
