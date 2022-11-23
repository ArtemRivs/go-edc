package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{7, 4, 5, 1, 2}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)
}
