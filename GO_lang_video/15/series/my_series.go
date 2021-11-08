package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")

}
func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 1000 {
		return nil, errors.New("n should be in [2, 1000]")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil

}
