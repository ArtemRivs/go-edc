package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := "abc"

	getUnicode(s)
}

func getUnicode(s string) string {

	r := []rune(s)

	fmt.Println("s:", s, "r:", r)

	var res string
	for _, ch := range s {
		if unicode.Is(unicode.Latin, ch) {
			res += string(ch)
		}
	}

	fmt.Println("res:", res)

	return res
}
