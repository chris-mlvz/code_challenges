package main

import (
	"fmt"
)

func main() {
	fmt.Println(MultiTable(5))
	fmt.Println(MultiTable(1))
}

func MultiTable(number int) string {
	var res string
	for i := 1; i <= 10; i++ {
		res += fmt.Sprintf("%v * %v = %v", i, number, i*number)
		if i != 10 {
			res += "\n"
		}
	}
	return res
}
