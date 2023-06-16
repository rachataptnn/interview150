package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println("res: ", removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}
