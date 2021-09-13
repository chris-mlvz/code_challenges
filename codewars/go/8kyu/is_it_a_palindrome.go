package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("a"))
	fmt.Println(IsPalindrome("aba"))
	fmt.Println(IsPalindrome("Abba"))
	fmt.Println(IsPalindrome("hello"))
}

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	reverseStr := Reverse(str)
	if reverseStr == str {
		return true
	}
	return false
}

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
