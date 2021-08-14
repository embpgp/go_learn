package main

import "fmt"

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen < 2 {
		return s
	}
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}
	for i := 0; i < sLen; i++ {
		dp[i][i] = true
	}
	resIdx := 0
	resLen := 1
	//遍历每一个字符串，从两边往中间找，采用动态规划思路
	for j := 1; j < sLen; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				//两边不相等直接放弃
				dp[i][j] = false
			} else {
				//中间有1个字符算OK
				if j-i < 3 {
					dp[i][j] = true
				} else {
					//否则继续比较中间的
					dp[i][j] = dp[i+1][j-1]
				}
			}
			//判断为回文串，且长度大于当前记录的最长的
			if dp[i][j] && j-i+1 > resLen {
				resIdx = i
				resLen = j - i + 1
			}

		}
	}
	return s[resIdx : resIdx+resLen]
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
	s = "a"
	fmt.Println(longestPalindrome(s))
	s = "ac"
	fmt.Println(longestPalindrome(s))
}
