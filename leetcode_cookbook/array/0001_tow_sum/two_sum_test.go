package tow_sum

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i //哈希存储下来，key为数组的val, val为数组idx
	}
	return nil
}

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 11, 5}
	target := 9
	fmt.Println(twoSum(nums, target))
}
