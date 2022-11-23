package main

import "fmt"

func increment() func() int {
	test := 0
	return func() int {
		test += 10
		return test
	}
}

func main() {

	inc := increment()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

}
