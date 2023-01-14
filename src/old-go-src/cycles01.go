package main

import (
	"fmt"
)

func main() {

	array := make([]int, 0, 10)

	// for i := 0; i < 10; i++ {
	// 	array = append(array, i*5)
	// }

	i := 0

	for i < 10 {
		array = append(array, i)
		i++
	}
	fmt.Println("array: ", array)
	fmt.Println("array len: ", len(array))

}
