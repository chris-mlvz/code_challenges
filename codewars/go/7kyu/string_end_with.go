package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("abcde", "cde"))
	fmt.Println(solution("abcde", "abc"))
	fmt.Println(solution("abcde", ""))
}

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
