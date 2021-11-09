package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1, s2)第二个参数不能直接使用slice，使用...或者直接是1,2,3
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
