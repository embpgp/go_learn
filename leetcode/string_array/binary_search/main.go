package main

import "fmt"

func searchInsert(nums []int, target int) int {

	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}

	return -1

}

func main(){
	array := []int{1,3,5,6}
	fmt.Println(searchInsert(array, 5))

}