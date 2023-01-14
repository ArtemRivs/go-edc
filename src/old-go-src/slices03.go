package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// array := make([]int, 32, 32)

	fmt.Println("array: ", array)
	fmt.Println("array len: ", len(array))

	slice01 := array[:3]
	fmt.Println("slice01 [:3] len: ", len(slice01))
	fmt.Println("slice01 [:3]: ", slice01)

	slice02 := array[3:7]
	fmt.Println("slice02 [3:7] len: ", len(slice02))
	fmt.Println("slice02 [3:7]: ", slice02)

	slice03 := array[5:]
	fmt.Println("slice03 [5:] len: ", len(slice03))
	fmt.Println("slice03 [5:]: ", slice03)

	slice04 := array[1:2]
	fmt.Println("slice04 [1:2] len: ", len(slice04))
	fmt.Println("slice04 [1:2]: ", slice04)

}
