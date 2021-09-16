package main

import (
	"fmt"
	"math"
)

func main() {
	solve(12, 20, 8)
}

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	tip := meal_cost * float64(tip_percent) / 100
	tax := meal_cost * float64(tax_percent) / 100
	total_cost := meal_cost + tip + tax
	fmt.Println(math.Round(total_cost))
}
