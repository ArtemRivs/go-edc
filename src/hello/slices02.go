package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println("Origin nums slice: ", nums)
	// Remove(nums, 5)
	// fmt.Println("Modified nums slice: ", nums)

	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	fmt.Println("nums slice: ", nums)
	newNums := nums[:3]
	fmt.Println("newNums slice [:3]: ", newNums)
	newNums = nums[2:]
	fmt.Println("newNums slice [2:]: ", newNums)
	newNums = nums[1:6]
	fmt.Println("newNums slice [1:6]: ", newNums)

	//Remove(nums, 1)
	//fmt.Println("Modified nums slice: ", nums)

}
func Remove(nums []string, i int) []string {

	// if i < 0 || i >= len(nums) {
	// 	return nums
	// }

	// var newNums []int
	// for j := 0; j < len(nums); j++ {
	// 	if j == i {
	// 		continue
	// 	}
	// 	newNums = append(newNums, nums[j])
	// }

	// nums = newNums

	// fmt.Println("newNums slice: ", newNums)

	// return nums

	//nums[i] = nums[len(nums)-1]

	//fmt.Println("nums slice: ", nums[:len(nums)-1])

	// nums = append(nums[:i], nums[i+1:]...)
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println("nums sliced: ", nums)

	return nums

}
