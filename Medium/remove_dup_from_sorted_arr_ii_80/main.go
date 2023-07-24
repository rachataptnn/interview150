package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	len, modNums := removeDuplicates(nums)
	fmt.Println("res: ", len, modNums)
}

func removeDuplicates(nums []int) (int, []int) {
	if len(nums) <= 2 {
		return len(nums), nums
	}

	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}

	return index, nums[:index]
}
