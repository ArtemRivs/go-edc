package main

import "fmt"

type Square struct {
	length float32
}
type Rectangle struct {
	firstSide  float32
	secondSide float32
}
type Circle struct {
	radius float32
}

type Shape interface {
	Area() float32
}

func (r Rectangle) Area() float32 {
	return r.firstSide * r.secondSide
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
	c := Rectangle{firstSide: 1.8, secondSide: 4.2}
	//площадь квадрата
	printShapeArea(a)
	//площадь окружности
	printShapeArea(b)
	//площадь прямоугольника
	printShapeArea(c)

}
