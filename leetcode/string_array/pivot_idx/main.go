package main

import "fmt"

func pivotIndex(nums []int) int {

	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] == (total - sum) {
			return i
		}
		sum += nums[i]
	}

	return -1
}

func main() {
	array := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(array))
}
