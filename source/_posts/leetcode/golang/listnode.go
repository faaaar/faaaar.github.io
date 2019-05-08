package leetcode

import (
	"strconv"
)

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 根据数字切片创建链表
func NewListNode(s []int) *ListNode {
	l := len(s)

	if l == 0 {
		return nil
	}

	node := &ListNode{
		Val: s[0],
	}
	head := node

	for i := 1; i < l; i++ {
		node.Next = &ListNode{
			Val: s[i],
		}
		node = node.Next
	}
	return head
}

// Compare 比较两个链表是否相等
func (l1 *ListNode) Compare(l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val == l2.Val {
			l1 = l1.Next
			l2 = l2.Next

			continue
		}

		return false
	}

	return l1 == nil && l2 == nil
}

// String 输出链表
func (l1 *ListNode) String() string {
	if l1 == nil {
		return ""
	}

	str := ""

	for l1 != nil {
		str += strconv.Itoa(l1.Val) + ","
		l1 = l1.Next
	}
	return str[:len(str)-1]
}
