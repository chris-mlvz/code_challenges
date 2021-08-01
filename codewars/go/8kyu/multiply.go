package main

import "fmt"

func main() {
	fmt.Println(Multiply(1, 1) == 1)
	fmt.Println(Multiply(1, 0) == 0)
	fmt.Println(Multiply(2, 5) == 10)
	fmt.Println(Multiply(5, 10) == 50)

}

func Multiply(a, b int) int {
	return a * b
}
