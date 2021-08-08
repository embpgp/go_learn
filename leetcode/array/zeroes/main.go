package main

import "fmt"

/*
 *  setZeroes 将矩阵为0的元素行和列均设置为0，思路：先便利一遍，找出对应的行和列，再次遍历清空即可。
 *  @Description:
 *  @param matrix
 */
func setZeroes(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	fmt.Println(matrix1)
	setZeroes(matrix1)
	fmt.Println(matrix1)

	fmt.Println(matrix2)
	setZeroes(matrix2)
	fmt.Println(matrix2)
}
