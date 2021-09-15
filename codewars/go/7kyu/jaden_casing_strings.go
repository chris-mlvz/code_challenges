package main

import "strings"

func main() {
	print(ToJadenCase("most trees are blue"))
}

func ToJadenCase(str string) string {
	return strings.Title(str)
}
