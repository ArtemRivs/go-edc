package main

// import "unicode"
import (
	"fmt"
	"reflect"
)

func main() {

	s := "✪ superstar ✪"
	fmt.Println("isASCII", isASCII(s))

}

func isASCII(s string) bool {

	res := true
	strSlice := []rune(s)
	minASCII := byte(0)
	maxASCII := byte(255)

	fmt.Println("minASCII:", reflect.TypeOf(minASCII), string(minASCII))
	fmt.Println("maxASCII:", string(maxASCII))
	fmt.Println("strSlice:", reflect.TypeOf(strSlice), strSlice)

	for _, ch := range strSlice {
		fmt.Println("ch:", ch, string(ch))

		if ch < 0 || ch > 255 {
			res = false
			break
		}
	}
	return res

}
