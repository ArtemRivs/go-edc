package main

import "fmt"

func main() {

// 	nums := make([]int, 5, 5)
// 	fmt.Println("nums: ", nums)
// 	fmt.Println("nums lenght: ", nums)

// 	nums2 := make([]int, 0, 10)
// 	fmt.Println("nums2: ", nums2)
// 	fmt.Println("nums2 lenght: ", len(nums2))

// 	nums2 = append(nums2, 100)
// 	nums2 = append(nums2, 200)
// 	fmt.Println("nums2: ", nums2)
// 	fmt.Println("nums2 lenght: ", len(nums2))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7}
	//mySlice := make([]int, 7, 7)

	mySlice = {1, 2, 3, 4, 5, 6, 7}
	
	fmt.Println("len: ", len(mySlice), "cap: ", cap(mySlice))
	fmt.Println("mySlice: ", mySlice)	
	
	
	
}
