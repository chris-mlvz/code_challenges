package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"
	print(PrinterError(s))
}

func PrinterError(s string) string {
	allowedLetters := "abcdefghijklm"
	errorsCounts := 0
	for _, letter := range s {
		if !strings.Contains(allowedLetters, string(letter)) {
			errorsCounts++
		}
	}
	return fmt.Sprintf("%v/%v", errorsCounts, len([]rune(s)))
}
