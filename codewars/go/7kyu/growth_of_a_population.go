package main

import "fmt"

func main() {
	fmt.Println(NbYear(1000, 2.0, 50, 1214))
	fmt.Println(NbYear(1500000, 0.0, 10000, 2000000))
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	for p0 < p {
		p0 = p0 + int(float64(p0)*percent/100) + aug
		years++
	}
	return years
}
