package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{7, 4, 5, 1, 2}
	fmt.Println("nums", nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println("nums", nums)

	nums2 := []int{8, 4, 1, 4, 5, 4, 9}
	fmt.Println("nums2", nums2)
	sort.SliceStable(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	fmt.Println("nums2", nums2)

}
