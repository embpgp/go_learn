package main

import "fmt"

/*
 *  findDiagonalOrder
 *  @Description: 对角线遍历
 *  @param mat
 *  @return []int
 */
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	c := 0
	r := 0
	ret := make([]int, 0)
	for i := 0; i < m+n-1; i++ {
		//左下->右上
		if i%2 == 0 {
			for j := r; j >= i-c; j-- {
				ret = append(ret, mat[j][i-j])
			}
		} else { //右上->左下
			for j := c; j >= i-r; j-- {
				ret = append(ret, mat[i-j][j])
			}
		}
		if r >= m-1 {
			r = m - 1
		} else {
			r = r + 1
		}
		if c >= n-1 {
			c = n - 1
		} else {
			c = c + 1
		}

	}
	return ret
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	fmt.Println(findDiagonalOrder(matrix))
}
