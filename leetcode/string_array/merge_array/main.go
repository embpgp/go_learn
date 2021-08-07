package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a

}

type Element [][]int

func (p Element) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Element) Len() int      { return len(p) }

//判断i位置的元素是否比j位置的元素小，如果为真，sort方法会调用swap函数交换
//但是sort函数本身实现是根据从后往前判断的，因此还是默认为升序
func (p Element) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func merge(intervals [][]int) [][]int {

	ret := make([][]int, 0)
	sort.Sort(Element(intervals))
	retIndex := 0
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if ret[retIndex][1] < intervals[i][0] { // 当前扫描区间和前一区间不重合，则直接添加到结果中
			ret = append(ret, intervals[i])
			retIndex++
		} else { // 相邻两个区间重合，则合并成一个区间
			ret[retIndex][0] = min(ret[retIndex][0], intervals[i][0])
			ret[retIndex][1] = max(ret[retIndex][1], intervals[i][1])
		}
	}

	return ret
}

func main() {
	array := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	out := merge(array)
	fmt.Println((out))
}
