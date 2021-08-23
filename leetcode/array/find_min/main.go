package main

import "fmt"

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r-l)/2 + l
		// 中值>右值, 中值肯定不是最小值
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			// 中值<左值, 中值有可能是最小值
			r = mid
		}
	}
	return nums[l]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums))
	nums = []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}
