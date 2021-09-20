package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	var words []string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var tempStr string
		fmt.Scan(&tempStr)
		words = append(words, tempStr)
	}
	for _, word := range words {
		var firstStr, secondStr string
		for i, v := range word {
			if i%2 == 0 {
				firstStr += string(v)
			} else {
				secondStr += string(v)
			}
		}
		fmt.Println(firstStr, secondStr)
	}
}
