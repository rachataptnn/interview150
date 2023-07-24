package main

func main() {
	nums := []int{3, 2, 2, 3}
	num := 3
	removeElement(nums, num)
}

func removeElement(nums []int, val int) int {
	count := 0
	for _, num := range nums {
		if num != val {
			nums[count] = num
			count++
		}
	}
	return count
}
