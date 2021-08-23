package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	res := []int{}
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			res = append(res, l+1, r+1)
			break
		}
	}
	return res
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
	numbers = []int{2, 3, 4}
	target = 6
	fmt.Println(twoSum(numbers, target))
	numbers = []int{-1, 0}
	target = -1
	fmt.Println(twoSum(numbers, target))

}
