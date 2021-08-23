package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] == val {
			fast++
			continue
		}
		nums[slow] = nums[fast]
		slow++
		fast++

	}
	return slow
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElement(nums, val))
}
