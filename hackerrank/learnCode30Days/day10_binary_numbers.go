package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)
	s := fmt.Sprintf("%b", n)
	tempCount := 0
	counts := []int{}
	for _, v := range s {
		if v == '0' {
			counts = append(counts, tempCount)
			tempCount = 0
		} else {
			tempCount++
		}
	}
	counts = append(counts, tempCount)
	sort.Ints(counts)
	fmt.Println(n, s, counts, counts[len(counts)-1])
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
