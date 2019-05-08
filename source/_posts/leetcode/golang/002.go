package leetcode

// 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// AddTwoNumbers002 两数相加
func AddTwoNumbers002(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	for l1 != nil || l2 != nil {
		l1.Val += l2.Val
		l3 := l1
		for l3.Val >= 10 {
			l3.Val -= 10

			if l3.Next != nil {
				l3.Next.Val++
				l3 = l3.Next
			} else {
				l3.Next = &ListNode{1, nil}
			}
		}

		if l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			break
		}

		if l1.Next != nil && l2.Next == nil {
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return head
}

// AddTwoNumbersRecursive002 两数相加递归
func AddTwoNumbersRecursive002(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return l2
	}

	if l1 != nil && l2 == nil {
		return l1
	}

	l1.Val += l2.Val
	recursive002(l1)

	l1.Next = AddTwoNumbersRecursive002(l1.Next, l2.Next)

	return l1
}

// recursive002 递归部分
func recursive002(l *ListNode) {
	if l.Val >= 10 {
		l.Val -= 10

		if l.Next != nil {
			l.Next.Val++
			recursive002(l.Next)
		} else {
			l.Next = &ListNode{1, nil}
		}
	}
}
