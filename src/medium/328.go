package main

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	pl, p := head, head.Next

	for p != nil && p.Next != nil {
		p3 := p
		p = p.Next
		p2 := pl.Next
		p3.Next = p.Next
		p.Next = p2
		pl.Next = p
		pl = p
		p = p3.Next
	}

	return head
}
