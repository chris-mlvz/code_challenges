package main

import "fmt"

func main() {
	fmt.Println(PositiveSum([]int{-1, 2, 3, 4, -5}) == 9)
}

func PositiveSum(numbers []int) int {
	var total int
	for _, v := range numbers {
		if v > 0 {
			total += v
		}

	}
	return total
}
