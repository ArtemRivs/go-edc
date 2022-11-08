package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	// tom := person{name: "Tom", age: 23}
	// fmt.Println("tom:", tom)

	// var tomPointer *person = &tom
	// tomPointer.age = 40

	// fmt.Println("tom:", tom)
	// fmt.Println("tomPointer:", tomPointer)
	// fmt.Println("*(tomPointer):", *(tomPointer))
	// fmt.Println("*tomPointer:", *tomPointer)

	x := 10
	y := &x
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("&x:", &x)
	fmt.Println("*y:", *y)
	*y = 88
	fmt.Println("x:", x)

}
