package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	head := &Node{Val: 1}
	cur := head
	for i := 2; i < 12; i++ {
		new := &Node{Val: i}
		cur.Next = new
		cur = new
	}
	printList(head)
	head = swap(head)
	printList(head)
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Printf(" -> ")
		}
		head = head.Next
	}
	fmt.Printf("\n")
}

func swap(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = newHead.Next
	newHead.Next = head

	// Swap the rest
	head.Next = swap(head.Next)
	return newHead
}
