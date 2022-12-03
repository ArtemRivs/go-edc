package main

import "fmt"

type Shape interface {
	getArea() float32
}
type Circle struct {
	radius float32
}
type Square struct {
	length float32
}

func (c Circle) getArea() float32 {
	return c.radius * c.radius * 3.14
}

func (s Square) getArea() float32 {
	return s.length * s.length
}

func main() {

	var S1 Shape = Circle{radius: 200.5}
	var S2 Shape = Square{length: 9999999.5}

	// fmt.Println(S1)
	// fmt.Println(S2)
	fmt.Println("Circle area: ", S1.getArea())
	fmt.Println("Square area: ", S2.getArea())

}
