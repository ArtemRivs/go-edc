package main

import "fmt"

func main() {
  nums[]int := {1, 2, 3, 4, 5, 6, 7, 8}
  fmt.Println("Origin nums slice: ", nums)
  Remove(nums, 5)
  fmt.Println("Modified nums slice: ", nums)
  
}
func Remove(nums []int, i int) []int {

	if i < 0 || i >= len(nums) {
		return nums
	}

	var newNums[]int
	for (j := 0; j < len(nums); j++) {
		if j == i {
			continue
		}
		append(newNums, nums[j])
	}

	nums = newNums

	return nums

}
