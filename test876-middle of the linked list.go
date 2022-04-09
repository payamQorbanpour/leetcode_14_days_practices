package main

import "fmt"

type ListNode struct {
	Val int 
	Next *ListNode
}

func findLength(list *ListNode) int {
	var temp ListNode = *list.Next
	var count int = 0
	for (*temp.Next != nil) {
		temp = *temp.Next

		count++
	}
	return count
}

// func middleNode(head *ListNode) *ListNode {
// 	for i := 0; i < len(head); i++ {

// 	}

// 	return head
// }

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	
	fmt.Println(findLength(&head))
	// fmt.Println(middleNode(&head))
}