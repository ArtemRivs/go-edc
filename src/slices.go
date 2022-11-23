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

	mySlice := []int{1, 2, 3, 4, 5, 6}

	mySlice2 := mySlice

	fmt.Println("mySlice: ", mySlice)
	fmt.Println("len: ", len(mySlice), "cap: ", cap(mySlice))

	mySlice[0] = 100
	mySlice[5] = 600

	fmt.Println("mySlice2: ", mySlice2)
	fmt.Println("len: ", len(mySlice2), "cap: ", cap(mySlice2))
	
	aSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("cap:", cap(aSlice), "len:", len(aSlice))
	fmt.Println(aSlice)

	change(&aSlice)

	fmt.Println("cap:", cap(aSlice), "len:", len(aSlice))
	fmt.Println(aSlice)

	dest := make([]int, 3)
	dest = []int{1, 2, 3}
	fmt.Println("dest", dest)

	var nums = []int{}
	nums = []int{1, 2, 3}
	fmt.Println("nums", nums)
	
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s:", s)

	//s = s[:len(s)-1]
	//fmt.Println("s:", s)

	s = s[1:]
	fmt.Println("s:", s)

	i := 3
	s = append(s[:i], s[i+1:]...)
	fmt.Println("s:", s)
	
	
	
}

func change(sl *[]int) {

	*sl = append(*sl, 127)

}
