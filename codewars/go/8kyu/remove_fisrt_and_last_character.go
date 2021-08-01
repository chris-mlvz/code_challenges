package main

import "fmt"

func main() {
	fmt.Println(RemoveChar("person") == "erso")
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
