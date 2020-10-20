package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	nodeList := make([]*ListNode, 0, 100)
	for p := head; p != nil; p = p.Next {
		nodeList = append(nodeList, p)
	}

	i, j := 0, len(nodeList)-1
	for i < j-1 {
		nodeList[i].Next = nodeList[j]
		nodeList[j].Next = nodeList[i+1]
		nodeList[j-1].Next = nil
		i++
		j--
	}
}
