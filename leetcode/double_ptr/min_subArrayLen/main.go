package main

import "fmt"

/*
 *  minSubArrayLen
 *  @Description: 采用滑动窗口左右两个指针查找最短>=target值的子串
 *  @param target
 *  @param nums
 *  @return int
 */
func minSubArrayLen(target int, nums []int) int {
	res := 1<<31 - 1

	l, sum := 0, 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]

		for sum >= target {
			if r-l+1 < res {
				res = r - l + 1
			}
			sum -= nums[l] // 收缩左边界
			l++
		}
	}
	if res == 1<<31-1 {
		res = 0
	}

	return res

}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
	nums = []int{1, 4, 4}
	target = 4
	fmt.Println(minSubArrayLen(target, nums))
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	target = 11
	fmt.Println(minSubArrayLen(target, nums))
}
