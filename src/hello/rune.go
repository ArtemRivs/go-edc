package main

// import "unicode"
import "fmt"

func main() {

	s := "Hello!"
	fmt.Println("isASCII", isASCII(s))

}

func isASCII(s string) bool {

	res := true
	strSlice := []rune(s)

	fmt.Println("strSlice:", strSlice)

	for _, ch := range strSlice {
		fmt.Println("ch:", ch, string(ch))

		if ch < 0 || ch > 255 {
			res = false
			break
		}
	}
	return res

}
