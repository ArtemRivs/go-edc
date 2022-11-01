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

	// slice := make([]string, 0, 6)
	slice := []string{"one", "two", "three", "four", "five", "six"}

	// for i, value := range slice {
	// 	fmt.Println("index: ", i, "value: ", value)
	// }
	// fmt.Println("slice 01: ", slice)
	// slice = someChanges(slice)
	// fmt.Println("slice 03: ", slice)
	// fmt.Println("slice: ", slice)

	_slice := slice

	fmt.Println("slice  01: ", slice)
	fmt.Println("_slice 01: ", _slice)

	_slice[0] = "ten"
	_slice[len(_slice)-1] = "twenty"

	fmt.Println("slice:  02", slice)
	fmt.Println("_slice: 02", _slice)

	nums := []int{1, 2, 3, 4, 5}
	numsCp := make([]int, 0)
	copy(numsCp, nums)
	fmt.Println("nums: ", nums)
	fmt.Println("numsCp: ", numsCp)

	numsB := make([]int, len(nums))
	copy(numsB, nums)
	fmt.Println("numsB: ", numsB)

}

func someChanges(s []string) []string {

	for i, value := range s {
		s[i] = value + "_01"
	}

	fmt.Println("slice 02: ", s)

	return s
}
