package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = nums1[:m]
	nums2 = nums2[:n]

	arr3 := append(nums1, nums2...)

	sort.Ints(arr3)

	fmt.Println(arr3)
	return arr3
}
