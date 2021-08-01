package main

import (
	"fmt"
)

func main() {
	fmt.Println(Opposite(1) == -1)
}

func Opposite(value int) int {
	return -value
}
