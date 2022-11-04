package solution

import (
	"fmt"
	"reflect"
)

func main() {

	nextASCII(byte('a'))
	prevASCII(byte('o'))
}

func nextASCII(b byte) byte {
	fmt.Println(reflect.TypeOf(b))
	return b
}

func prevASCII(b byte) byte {
	fmt.Println(reflect.TypeOf(b))
	return b
}
