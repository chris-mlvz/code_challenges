package main

import "fmt"

func main() {
	fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
	fmt.Println(multipleOfIndex([]int{68, -1, 1, -7, 10, 10}))
	fmt.Println(multipleOfIndex([]int{11, -11}))
	fmt.Println(multipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}))
	fmt.Println(multipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51}))
	fmt.Println(multipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35}))
}

func multipleOfIndex(ints []int) []int {
	indexSlice := []int{}
	for i, v := range ints {
		if i == 0 {
			continue
		}
		if v%i == 0 {
			indexSlice = append(indexSlice, v)
		}
	}
	return indexSlice
}
