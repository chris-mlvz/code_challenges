package main

import "fmt"

func main() {
	fmt.Println(EvenOrOdd(13) == "Odd")
	fmt.Println(EvenOrOdd(6) == "Even")
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
