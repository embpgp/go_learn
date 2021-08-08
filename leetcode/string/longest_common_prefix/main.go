package main

import "fmt"

/*
 *  longestCommonPrefix
 *  @Description:编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
 *  @param strs
 *  @return string
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var ret string
	var flag int
	for i := 0; i < len(strs[0]); i++ {
		flag = 0
		//以第一个元素为基准，根据字符串数组的长度遍历，以此对比每一个字符
		for j := 0; j < len(strs); j++ {
			//注意i不能越界
			if i < len(strs[j]) && strs[0][i] == strs[j][i] {
				flag++
			} else {
				flag = -1
				break
			}

		}
		//每个字符串都有这个字符，就追加这个字符
		if flag == len(strs) {
			ret = ret + string(strs[0][i])
		} else {
			break
		}
	}
	return ret
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
