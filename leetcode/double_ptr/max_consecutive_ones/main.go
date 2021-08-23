package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	slow, fast := -1, 0
	cnt := 0
	for ; fast < len(nums)+1; fast++ {
		//处理最后一个元素
		if fast == len(nums) || nums[fast] != 1 {
			if fast-slow-1 > cnt {
				cnt = fast - slow - 1
			}
			slow = fast
		}
	}
	return cnt
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
