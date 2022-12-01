package main

import "fmt"

type Square struct {
	length float32
}

type Circle struct {
	radius float32
}

type Shape interface {
	Area() float32
}

func (l Square) Area() float32 {
	return l.length * l.length
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * 3.14
}

func printShapeArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {

	fmt.Println("GO!")
	a := Square{length: 15.5}
	b := Circle{radius: 10.4}
	printShapeArea(a)
	printShapeArea(b)
}
