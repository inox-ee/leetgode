package problem0002

func _addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// var head *ListNode // nil ポインタで初期化されてしまうので注意
	var head = &ListNode{}
	var cur = head
	var c0 = 0

	for l1 != nil || l2 != nil || c0 > 0 {
		sum := c0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		c0 = sum / 10
		cur = cur.Next
	}

	return head.Next
}
