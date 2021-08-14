package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	//割切字符数组，这里其实开辟了新的空间
	str := strings.Split(s, " ")
	sLen := len(str)
	n := 0
	i := 0
	for {
		if i >= sLen {
			break
		}
		if str[i] == "" {
			i++
			continue
		}

		str[n], str[i] = str[i], str[n]
		n++
		i++

	}

	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}

	return strings.Join(str[:n], " ")

}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
	s = "  hello world  "
	fmt.Println(reverseWords(s))
	s = "a good   example"
	fmt.Println(reverseWords(s))
	s = "  Bob    Loves  Alice   "
	fmt.Println(reverseWords(s))
	s = "Alice does not even like bob"
	fmt.Println(reverseWords(s))
}
