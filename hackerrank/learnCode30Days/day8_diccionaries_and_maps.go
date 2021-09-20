package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	myMap := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	n, err := strconv.Atoi(scanLine(scanner))
	checkError(err)

	for i := 0; i < n; i++ {
		strTemp := scanLine(scanner)
		sliceTemp := strings.Split(strTemp, " ")
		myMap[sliceTemp[0]] = sliceTemp[1]
	}
	for {
		strTemp := scanLine(scanner)
		if strTemp == "" {
			break
		}
		if _, ok := myMap[strTemp]; ok {
			fmt.Printf("%v=%v\n", strTemp, myMap[strTemp])
		} else {
			fmt.Println("Not found")
		}
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scanLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
