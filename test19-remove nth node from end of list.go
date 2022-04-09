package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	n := 2

	fmt.Println(removeNthFromEnd(&head, n))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := head
	length := 1
	for head.Next != nil {
		length++
		head = head.Next
	}

	if length == 1 {
		head = nil
		return head
	}

	previousNodeIndex := length - n

	if previousNodeIndex == 0 {
		tmp = tmp.Next
		return tmp
	}
	count := 1
	tmp2 := tmp
	for tmp2.Next != nil {
		if count == previousNodeIndex {
			tmp2.Next = tmp2.Next.Next
			return tmp
		}
		count++
		tmp2 = tmp2.Next
	}

	return tmp
}
