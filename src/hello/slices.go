package main

import "fmt"

func main() {

	// s := []int{10, 20, 30}
	// fmt.Println("s: ", s)
	//fmt.Println("s[1]: ", s[0])
	// fmt.Println("s[2]: ", s[1])
	// fmt.Println("s[3]: ", s[2])
	// s[2] = 40
	// s = append(s, 50)
	// fmt.Println("s append 3: ", s)
	// s[3] = 60
	// fmt.Println("s change 3: ", s)

	nums := make([]int, 5, 5)
	fmt.Println("nums: ", nums)
	fmt.Println("nums lenght: ", nums)

	nums2 := make([]int, 0, 10)
	fmt.Println("nums2: ", nums2)
	fmt.Println("nums2 lenght: ", len(nums2))

	// nums2[0] = 100
	nums2 = append(nums2, 100)
	nums2 = append(nums2, 200)
	fmt.Println("nums2: ", nums2)
	fmt.Println("nums2 lenght: ", len(nums2))

}
