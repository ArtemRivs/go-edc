package solution

func SafeWrite(nums [5]int, i, val int) [5]int {

	var arrМахIndex = len(nums) - 1;

	if i >= 0 && i <= arrМахIndex {
		nums[i] = val
	}

	return nums

}
