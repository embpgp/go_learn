package main

import "fmt"

func rotate(matrix [][]int) {
	//沿着左上到右下的对角线翻转
	for row := 0; row < len(matrix); row++ {
		for col := 0; col <= row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	//沿着1/2列的地方水平翻转
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix)/2; col++ {
			matrix[row][col], matrix[row][len(matrix)-col-1] = matrix[row][len(matrix)-col-1], matrix[row][col]
		}
	}
}

func main() {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	fmt.Println(matrix1)
	rotate(matrix1)
	fmt.Println(matrix1)

	fmt.Println(matrix2)
	rotate(matrix2)
	fmt.Println(matrix2)

}
