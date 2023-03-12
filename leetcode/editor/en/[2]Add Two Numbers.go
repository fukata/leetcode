//You are given two non-empty linked lists representing two non-negative 
//integers. The digits are stored in reverse order, and each of their nodes contains a 
//single digit. Add the two numbers and return the sum as a linked list. 
//
// You may assume the two numbers do not contain any leading zero, except the 
//number 0 itself. 
//
// 
// Example 1: 
// 
// 
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
// 
//
// Example 2: 
//
// 
//Input: l1 = [0], l2 = [0]
//Output: [0]
// 
//
// Example 3: 
//
// 
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in each linked list is in the range [1, 100]. 
// 0 <= Node.val <= 9 
// It is guaranteed that the list represents a number that does not have 
//leading zeros. 
// 
//
// Related Topics Linked List Math Recursion 👍 24798 👎 4794


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	node := result
	node1 := l1
	node2 := l2
	next_value := 1
	for {
		v := node.Val
		if node1 != nil {
			v += node1.Val
		}
		if node2 != nil {
			v += node2.Val
		}
		node.Val = v % 10

		if v >= 10 {
			next_value = 1
		} else {
			next_value = 0
		}

		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
		if node1 == nil && node2 == nil {
			if next_value > 0 {
				node.Next = &ListNode{next_value, nil}
			}
			break
		}

		next_node := &ListNode{next_value, nil}
		node.Next = next_node
		node = next_node
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
