package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	slow, fast := 0, 1
	nLen := len(nums)
	for ; fast < nLen; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]

		}
	}
	return slow + 1
	/*

		n := len(nums)
		if n == 0 {
			return 0
		}
		slow := 1
		for fast := 1; fast < n; fast++ {
			if nums[fast] != nums[fast-1] {
				nums[slow] = nums[fast]
				slow++
			}
		}
		return slow
	*/

}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
