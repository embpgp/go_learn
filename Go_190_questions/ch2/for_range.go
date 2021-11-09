package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val //val是一个临时拷贝的变量，所以相当于go运行时最后一次仍然使用这个地址存储3,m存的地址都是一样的
		fmt.Printf("0x%x, %d\n", &val, val)
		/*结果如下：
		0xc00000a0d8, 0
		0xc00000a0d8, 1
		0xc00000a0d8, 2
		0xc00000a0d8, 3
		3 -> 3
		0 -> 3
		1 -> 3
		2 -> 3
		*/
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
