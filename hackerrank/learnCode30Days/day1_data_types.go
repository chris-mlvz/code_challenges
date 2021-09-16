package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var myInt int
	var myDouble float64
	var myString string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	myInt, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	myDouble, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	myString = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(myInt + int(i))
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", myDouble+d)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + myString)
}
