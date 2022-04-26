package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	fmt.Println(connect(root))
}

func connect(root *Node) *Node {
	populate(root)

	return root
}

func populate(node *Node) {
	if node == nil {
		return
	}
	if node.Left == nil {
		return
	}

	node.Left.Next = node.Right
	if node.Next != nil {
		node.Right.Next = node.Next.Left
	}
	populate(node.Left)
	populate(node.Right)
}
