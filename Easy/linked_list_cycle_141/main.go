package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{3, 2, 0, -4}
	pos := 1

	cycleLL := createCyclicLinkedList(nums, pos)
	printCyclicLinkedList(cycleLL)
}

// helper func from chatGPT <3 ty so much
func createCyclicLinkedList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(values))

	// Create nodes for the linked list
	for i, val := range values {
		nodes[i] = &ListNode{Val: val}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	// Link the last node to the node at the given position to create a cycle
	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
	}

	return nodes[0]
}

// helper func
func printCyclicLinkedList(head *ListNode) {
	visited := make(map[*ListNode]bool)

	for head != nil {
		if visited[head] {
			fmt.Println("Cycle Detected. Stopping the printing.")
			break
		}
		visited[head] = true
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}
