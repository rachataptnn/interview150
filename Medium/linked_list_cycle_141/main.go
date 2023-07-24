package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	nums := []int{3, 2, 0, -4}
	// pos := 1
	linkedList := createLinkedList(nums)
	printLinkedList(linkedList)
}

// helper func
func createLinkedList(input []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, val := range input {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// helper func
func printLinkedList(input *ListNode) {
	for input != nil {
		fmt.Print(input.Val, " -> ")
		input = input.Next
	}
	fmt.Println("nil")
}

func hasCycle(head *ListNode) bool {
	return false
}
