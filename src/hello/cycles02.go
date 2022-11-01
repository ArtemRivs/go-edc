package main

// func Map(strs []string, mapFunc func(s string) string) []string {
// 	slice := make([]string, len(strs))
// 	for i, str := range strs {
//     	slice[i] = mapFunc(str)
// 	}
// 	return slice
// }

import "fmt"

func main() {

	// slice := []string{"one", "two", "three", "four", "five", "six"}
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println("index: ", i, "value: ", slice[i])

	// }

	slice := make([]string, 0, 6)
	slice = append(slice, "one")
	slice = append(slice, "two")
	slice = append(slice, "three")
	slice = append(slice, "four")
	slice = append(slice, "five")
	slice = append(slice, "six")
	fmt.Println("slice: ", slice)

}
