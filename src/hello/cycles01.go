package main

import (
	"fmt"
)

func main() {

	array := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		array = append(array, i*5)
	}
	// array := make([]int, 32, 32)

	fmt.Println("array: ", array)
	fmt.Println("array len: ", len(array))

}
