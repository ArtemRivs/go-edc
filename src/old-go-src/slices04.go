package solution

func IntsCopy(src []int, maxLen int) []int {

	srcCopy := []int{}

	if maxLen > len(src) {
		srcCopy = make([]int,len(src))
		copy(srcCopy, src)
	} else if maxLen > 0 {
		srcCopy = make([]int, maxLen)
		copy(srcCopy, src)
	}

	return srcCopy

}
