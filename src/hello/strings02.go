package main

import (
	"fmt"
	"reflect"
)

func main() {

	nextASCII(byte('a'))
	prevASCII(byte('o'))
}

func nextASCII(b byte) byte {

	b++
	fmt.Println("Type of b", reflect.TypeOf(b), "Value:", b)

	return b
}

func prevASCII(b byte) byte {

	b--
	fmt.Println("Type of b", reflect.TypeOf(b), "Value:", b)
	return b
}
