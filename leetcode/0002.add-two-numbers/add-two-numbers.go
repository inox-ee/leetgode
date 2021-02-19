package problem0002

// ListNode is a definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func divmod(i1 int, i2 int) (int, int) {
	return i1 / i2, i1 % i2
}

func calcTwoListNode(l1 *ListNode, l2 *ListNode, c0 int) *ListNode {
	if l1 == nil && l2 == nil {
		if c0 == 0 {
			return nil
		}
		return &ListNode{Val: c0, Next: nil}
	}
	if l1 == nil {
		l1 = &ListNode{Val: 0, Next: nil}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0, Next: nil}
	}
	var result *ListNode
	cc, num := divmod(l1.Val+l2.Val+c0, 10)
	var nextLN *ListNode
	nextLN = calcTwoListNode(l1.Next, l2.Next, cc)
	result = &ListNode{Val: num, Next: nextLN}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return calcTwoListNode(l1, l2, 0)
}

/*
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{} // 先頭を保持
	cur := result         // for 文中で値を更新
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 { // 終了条件は、「l1 が nil」 かつ 「l2 が nil」 かつ 「carry が 0」
		sum := carry

		if l1 != nil { // 「0 を足す」のではなく「何もしない」のが楽
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil { // 「0 を足す」のではなく「何もしない」のが楽
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		cur = cur.Next
	}

	return result.Next
}
*/
