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

	// fmt.Println("nums2", nums2)
	sort.SliceStable(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	// fmt.Println("nums2", nums2)
	fmt.Println("nums2", nums2)

	c := 0
	len_nums := len(nums2)
	for c < len_nums {
		// fmt.Println("c: ", c, "nums2[c]: ", nums2[c], "nums2[c+1]: ", nums2[c+1])

		if nums2[c] == nums2[c+1] {
			nums2 = append(nums2[:c], nums2[c+1:]...)
			// fmt.Println("c: ", c, "nums2: ", nums2)
		}
		len_nums = len(nums2)
		// fmt.Println("len_nums: ", len_nums)
		c++

		if c+1 == len_nums {
			break
		}
	}
	fmt.Println("end nums2", nums2)

}
