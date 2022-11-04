package main

import "fmt"

func main() {
	// m := map[int]string{1: "one", 2: "two", 3: "three"}
	// fmt.Println("m: ", m)

	str := map[string]int{"one": 1, "two": 2, "three": 3}
	// fmt.Println("str: ", str)

	for c, value := range str {
		fmt.Println("c: ", c, "value: ", value)
	}
	// fmt.Println("str['one']: ", str["one"])
	// fmt.Println("len(str): ", len(str))
	fmt.Println("str['one']: ", str["_one"])

}