package main

import (
	"fmt"
	"reflect"
)

func shiftASCII(s string, step int) string {

	var str string
	for p:=0; p > len(s); p++ {

		b := int(s[p])
		b = b + step
		if b >256 {
			b = 256
		}
	
		str = str + string(b)
	}

	return str

}
