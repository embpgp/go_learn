package main

import "fmt"

func reverseString(s []byte) {
	sLen := len(s)
	i, j := 0, sLen-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))

}
