package main

import "fmt"

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	swap(head, countList(head)-1)
	return head
}

func countList(head *ListNode) int {
	c := 0
	for head != nil {
		head = head.Next
		c++
	}
	return c
}

func swap(head *ListNode, n int) {
	if n < 0 {
		return
	}
	fmt.Println("n", n)
	newHead := head
	for i := 0; i < n; i++ {
		newHead = newHead.Next
	}
	newHead.Val, head.Val = head.Val, newHead.Val
	swap(head.Next, n-2)
}

func main() {
	head := &ListNode{Val: 1}
	cur := head
	for i := 2; i <= 5; i++ {
		new := &ListNode{Val: i}
		cur.Next = new
		cur = new
	}
	printList(head)
	head = reverseList(head)
	printList(head)
}

func printList(head *ListNode) {
	i := 0
	for head != nil && i < 20 {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Printf(" -> ")
		}
		head = head.Next
		i++
	}
	fmt.Printf("\n")
}
