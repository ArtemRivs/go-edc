package main

import "fmt"

type Vehicle interface {
	move(int) int
}

func drive(v Vehicle, time int) int {
	return v.move(time)
}

type Car struct{ speed int }
type Aircraft struct{ speed int }

func (c Car) move(time int) int {
	return time * c.speed
}
func (a Aircraft) move(time int) int {
	return time * a.speed
}
func main() {
	lexus := Car{speed: 100}
	boing := Aircraft{speed: 1000}
	fmt.Println(drive(lexus, 6))
	fmt.Println(drive(boing, 5))
}
