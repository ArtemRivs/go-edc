package main

import (
	"fmt"
	"sort"
)

func main() {

	// nums := []int{7, 4, 5, 1, 2}
	// fmt.Println("nums", nums)
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	// fmt.Println("nums", nums)

	nums2 := []int64{8, 4, 1, 4, 6, 3, 4, 6, 9, 2}
	sort.SliceStable(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	fmt.Println("nums2", nums2)

	pointer := 0
	for c := 1; c < len(nums2); c++ {
		if nums2[pointer] != nums2[c] {
			pointer++
			nums2[pointer] = nums2[c]
		}
	}
	nums2 = nums2[:pointer+1]
	fmt.Println("pointer", pointer)
	fmt.Println("nums2", nums2)

}
