package main

import "fmt"

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	nLen := len(nums)
	for ; fast < nLen; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < nLen; slow++ {
		nums[slow] = 0
	}

	/* //非零值和零进行交换
	func moveZeroes(nums []int) {
	    left, right, n := 0, 0, len(nums)
	    for right < n {
	        if nums[right] != 0 {
	            nums[left], nums[right] = nums[right], nums[left]
	            left++
	        }
	        right++
	    }
	}
	*/
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
