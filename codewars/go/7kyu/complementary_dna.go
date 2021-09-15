package main

import (
	"fmt"
)

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}

var DNAMap = map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}

func DNAStrand(dna string) string {
	var newADN string
	for i := 0; i < len(dna); i++ {
		newADN += string(DNAMap[dna[i]])
	}
	return newADN
}
