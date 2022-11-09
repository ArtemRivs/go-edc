package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Printf("Now is: %v.%v.%v \n", now.Day(), now.Month(), now.Year())
	fmt.Println("year", year)
}
