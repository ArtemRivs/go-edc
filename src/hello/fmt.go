package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{Name: "Vasiliy", Age: 40}
	fmt.Println("simple:", p)
	fmt.Printf("detailed: %+v\n", p)
	fmt.Printf("golang: %#v\n", p)

}
