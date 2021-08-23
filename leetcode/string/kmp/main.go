package main

import "fmt"

/*
 *  getNext
 *  @Description:  获取next数组也是一个往回走的过程
 *  @param p
 *  @return []int
 */
func getNext(p string) []int {
	plen := len(p)
	next := make([]int, plen)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < plen-1 {
		if i == -1 || p[i] == p[j] {
			i++
			j++
			next[j] = i
		} else {
			i = next[i]
		}
	}
	return next
}

func KmpSearch(s, p string) int {
	i, j := 0, 0
	pLen := len(p)
	sLen := len(s)
	next := getNext(p)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] {
			j++
			i++

		} else { //跳回next数组的下标j位置,i继续往前走
			j = next[j]
		}
	}
	//如果j全部走完，说明匹配完成
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 1 { // 子串长度=1 时单独判断
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				break
			}
		}
		if i < len(haystack) {
			return i
		} else {
			return -1
		}
	}

	return KmpSearch(haystack, needle)
}

func main() {
	s := "hello"
	p := "ll"
	fmt.Println(strStr(s, p))
	s = "aaaaa	"
	p = "bba"
	fmt.Println(strStr(s, p))
	s = ""
	p = ""
	fmt.Println(strStr(s, p))
}
