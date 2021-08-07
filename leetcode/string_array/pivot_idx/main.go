package main

import "fmt"

/*
算法思路：
最终的状态，当找到找个节点的时候，其左边的和sum+num[i] = 右边的和(total-sum)
*/
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
