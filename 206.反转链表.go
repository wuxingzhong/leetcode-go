/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	// head   1 2 3 4 5 nil
	// behind nil->1->2->3->4->5

	var behind *ListNode
	for head != nil {
		// 保存next
		next := head.Next
		// 解开链, head加到behind前
		head.Next = behind
		behind = head
		// 移到next
		head = next
	}

	return behind
}

// @lc code=end

