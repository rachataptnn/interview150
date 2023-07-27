package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range s {
		m1[v] += 1
	}

	for _, v := range t {
		m2[v] += 1
	}

	if len(m1) == len(m2) {
		for _, v := range s {
			if m1[v] != m2[v] {
				return false
			}
		}
	} else {
		return false
	}

	return true
}
