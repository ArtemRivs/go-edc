package solution

// import "unicode"

func  isASCII(s string) bool{

	res := false
	strSlice := []rune(s)
	for _, ch := range strSlice {
		if ch < 0 || ch > 255 {
			res = true
			break
		}
	}
	return res

}
