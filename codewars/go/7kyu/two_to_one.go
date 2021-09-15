package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
	fmt.Println(TwoToOne("inmanylanguages", "theresapairoffunctions"))

}

func TwoToOne(s1 string, s2 string) string {
	allKeys := make(map[rune]bool)
	setSlice := []rune{}
	for _, letter := range s1 + s2 {
		if _, ok := allKeys[letter]; !ok {
			allKeys[letter] = true
			setSlice = append(setSlice, letter)
		}
	}
	sort.Slice(setSlice, func(i, j int) bool {
		return setSlice[i] < setSlice[j]
	})
	return string(setSlice)
}
