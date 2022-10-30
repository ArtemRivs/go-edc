package main

import (
	"fmt"
	"unsafe"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("array: ", array)
	fmt.Println("array size: ", unsafe.Sizeof(array))

}
