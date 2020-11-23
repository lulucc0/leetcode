package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	insertionSortList(head)
}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	p1, p2 := head, head.Next
	for p2 != nil {
		if p2.Val >= p1.Val {
			p1 = p2
			p2 = p2.Next
		} else {
			p1.Next = p2.Next
			p := p2
			p2 = p2.Next
			if p.Val < head.Val {
				p.Next = head
				head = p
			} else {
				pl, pr := head, head.Next
				for pr != nil {
					if pr.Val <= p.Val {
						pl.Next = p
						p.Next = pr
						break
					} else {
						pl = pr
						pr = pr.Next
					}
				}
			}
		}
	}

	return head
}
